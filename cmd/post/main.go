package main

import (
	"c8s/config"
	saws "c8s/internal/service/cloud/aws"
	"log"
)

func main() {
	conf := config.GetConfig()
	ssmClient, err := saws.NewSSMClient(conf)
	if err != nil {
		log.Fatalf("Failed to initialize SSM client: %v", err)
	}

	if err := ssmClient.KubeProxyClose(); err != nil {
		log.Fatalf("Failed to close KubeProxy: %v", err)
	}

}
