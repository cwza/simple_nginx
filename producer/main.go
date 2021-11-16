package main

import (
	"bytes"
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "cfgpath", "./producer.toml", "config file path")
}

func send(url string) (string, error) {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return "", fmt.Errorf("failed to create http request, %s", err)
	}

	client := http.Client{
		Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
		Timeout:   2 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to send http request, %s", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("http error, %s", resp.Status)
		}
		return "", fmt.Errorf("http error, %s::%s", resp.Status, string(respBody))
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read http response, %s", err)
	}
	return string(respBody), nil
}

func sends(url string, cnt int) {
	for i := 0; i < cnt; i++ {
		go func() {
			_, err := send(url)
			if err != nil {
				log.Printf("failed to send msg, %s", err)
			}
		}()
	}
}

func run(url string, genSecRateFunc func() int) {
	for range time.Tick(time.Second) {
		cnt := genSecRateFunc()
		sends(url, cnt)
		log.Printf("send %d msgs\n", cnt)
	}
}

func main() {
	flag.Parse()

	config, err := initConfig(configPath)
	if err != nil {
		log.Fatalf("failed to init config, %s", err)
	}
	log.Printf("config: %+v\n", config)

	genSecRateFunc := createGenSecRateFunc(createGenMinRateFunc(config.Rates, config.Cnts))
	run(config.ConsumerUrl, genSecRateFunc)
}
