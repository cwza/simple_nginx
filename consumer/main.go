package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "cfgpath", "./consumer.toml", "config file path")
}

func main() {
	flag.Parse()

	// init config
	config, err := initConfig(configPath)
	if err != nil {
		log.Fatalf("failed to init config, %s", err)
	}
	log.Printf("config: %+v\n", config)

	// create http server
	server := &http.Server{
		Addr: ":" + config.HttpPort,
	}
	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/cpu", createCpuFunc(config.Cpu.LoopCnt))
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	log.Printf("Server Started at port: %s", config.HttpPort)

	// graceful shutdown
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
	log.Print("Server Stopped")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(config.ShutdownTimeout)*time.Second)
	defer func() {
		log.Print("Shutdown Timeout")
		cancel()
	}()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
