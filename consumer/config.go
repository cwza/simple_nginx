package main

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	HttpPort string `mapstructure:"httpport"`
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

	return config, nil
}
