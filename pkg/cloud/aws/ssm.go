package aws

import (
	conf "c8s/config"

	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

func NewSSMClient(conf conf.Config) (*AWSClients, error) {
	cfg := cfg(conf)
	ssmClient := ssm.NewFromConfig(cfg)

	return &AWSClients{
		SSM: ssmClient,
	}, nil
}
