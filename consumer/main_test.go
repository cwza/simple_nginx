package main

import (
	"fmt"
	"os"
	"testing"
)

func TestInitConfig(t *testing.T) {
	os.Setenv("CPU_LOOPCNT", "5")
	config, err := initConfig("./consumer.toml")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("config: %+v\n", config)
}
