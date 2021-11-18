package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort        string `mapstructure:"httpport"`
	ShutdownTimeout int    `mapstructure:"shutdowntimeout"`
	Cpu             struct {
		LoopCnt int `mapstructure:"loopcnt"`
	} `mapstructure:"cpu"`
}

func initConfig(filepath string) (Config, error) {
	viper.SetConfigFile(filepath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))
	viper.AutomaticEnv()

	config := Config{}
	err := viper.ReadInConfig()
	if err != nil {
		return Config{}, fmt.Errorf("read config failed: %s", err)
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return Config{}, fmt.Errorf("unmarshal config failed: %s", err)
	}

	if config.HttpPort == "" {
		return Config{}, fmt.Errorf("config.HttpPort is invalid")
	}
	if config.ShutdownTimeout < 0 {
		return Config{}, fmt.Errorf("config.ShutdownTimeout is invalid")
	}

	if config.Cpu.LoopCnt < 0 {
		return Config{}, fmt.Errorf("config.Cpu.LoopCnt is invalid")
	}

	return config, nil
}
