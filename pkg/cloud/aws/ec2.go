package aws

import (
	conf "c8s/config"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
)

type Ec2InstanceInfo struct {
	InstanceID   string
	InstanceName string
	Status       string
	PrivateIP    string
	State        string
}

func NewEC2Client(conf conf.Config) (*AWSClients, error) {
	cfg := Cfg(conf)
	ec2Client := ec2.NewFromConfig(cfg)

	return &AWSClients{
		EC2: ec2Client,
	}, nil
}

func (c *AWSClients) ListInstances() []Ec2InstanceInfo {
	input := &ec2.DescribeInstancesInput{}
	result, err := c.EC2.DescribeInstances(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to describe instances, %v", err)
	}

	var instances []Ec2InstanceInfo
	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			if instance.PrivateIpAddress == nil {
				continue
			}

			instances = append(instances, Ec2InstanceInfo{
				InstanceID:   *instance.InstanceId,
				InstanceName: selectTagVal(instance.Tags, "Name"),
				Status:       string(instance.State.Name),
				PrivateIP:    string(*instance.PrivateIpAddress),
				State:        string(instance.State.Name),
			})

		}
	}
	return instances
}
