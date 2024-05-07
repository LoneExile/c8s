package service_aws

import (
	"c8s/config"
	svc "c8s/pkg/cloud/aws"
	"fmt"
)

func ListInstances() {
	conf := config.GetConfig()
	c, err := svc.NewEC2Client(conf)
	if err != nil {
		fmt.Println(err)
	}
	instances := c.ListInstances()

	instancesLen := len(instances)
	if instancesLen == 0 {
		fmt.Println("\nNo instances found.")
		return
	}

	fmt.Println("EC2 Instances:")
	for key, val := range instances {
		fmt.Printf(
			"%d: %s (%s) %s %s\n",
			key,
			val.InstanceName,
			val.Status,
			val.PrivateIP,
			val.State,
		)
	}
}
