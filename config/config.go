package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type KubeConfig struct {
	InsecureSkipVerify bool
}

type Config struct {
	Port       string
	KubeConfig KubeConfig
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetConfig() Config {

	initConfig()

	var C Config
	err := viper.Unmarshal(&C)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}
	return C
}
