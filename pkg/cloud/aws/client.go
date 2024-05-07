package aws

import (
	conf "c8s/config"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

// TODO: add default, check config file => ENV => ~/.aws/config
// - check if region is valid
var (
	Region  string
	Profile string
)

type AWSClients struct {
	EC2 *ec2.Client
	RDS *rds.Client
	SSM *ssm.Client
}

func cfg(conf conf.Config) aws.Config {
	profile := conf.AWS.Profile

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return cfg
}
