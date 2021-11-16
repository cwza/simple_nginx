package main

import (
	"flag"
	"log"
	"net/http"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "cfgpath", "./consumer.toml", "config file path")
}

func main() {
	flag.Parse()

	config, err := initConfig(configPath)
	if err != nil {
		log.Fatalf("failed to init config, %s", err)
	}
	log.Printf("config: %+v\n", config)

	http.HandleFunc("/health", healthCheck)
	http.HandleFunc("/cpu", createCpuFunc(config.Cpu.LoopCnt))

	log.Fatal(http.ListenAndServe(":"+config.HttpPort, nil))
}
