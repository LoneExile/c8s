package ssm

import (
	conf "c8s/config"
	aCfg "c8s/pkg/cloud/aws"
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
)

type AWSClients struct {
	SSM *ssm.Client
}

func NewSSMClient(conf conf.Config) (*AWSClients, error) {
	cfg := aCfg.Cfg(conf)
	ssmClient := ssm.NewFromConfig(cfg)

	return &AWSClients{
		SSM: ssmClient,
	}, nil
}

func (c *AWSClients) SendCommand(commands []string, instanceIDs []string) error {
	parameters := map[string][]string{
		"commands": commands,
	}
	input := &ssm.SendCommandInput{
		DocumentName:   aws.String("AWS-RunShellScript"),
		InstanceIds:    instanceIDs,
		Parameters:     parameters,
		TimeoutSeconds: aws.Int32(60),
	}

	result, err := c.SSM.SendCommand(context.TODO(), input)
	if err != nil {
		log.Fatalf("Error sending command to SSM: %v", err)
		return err
	}

	log.Printf("Command sent to SSM: %v", result.Command.CommandId)
	return nil
}

func (c *AWSClients) ForwordPort(instanceID string, localPort int, remotePort int, conf conf.Config) error {
	cfg := aCfg.Cfg(conf)

	in := PortForwardingInput{
		Target:     instanceID,
		RemotePort: remotePort,
		LocalPort:  localPort,
	}

	err := PortPluginSession(cfg, &in)
	if err != nil {
		return err
	}

	return nil
}
