package config

import (
	"log"

	"github.com/spf13/viper"
)

type KubeConfig struct {
	ConfigPath         string
	ProxyUrl           string
	InsecureSkipVerify bool
}

type Application struct {
	Port string
}

type AWS struct {
	Profile    string
	InstanceID string
	LocalPort  int
	RemotePort int
}

type Config struct {
	App        Application
	KubeConfig KubeConfig
	AWS        AWS
}

func initConfig() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("fatal error config file: %v", err)
	}
}

func GetConfig() Config {

	initConfig()

	var C Config
	err := viper.Unmarshal(&C)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	return C
}
