package service_aws

import (
	"c8s/config"
	"c8s/pkg/cloud/aws/ssm"
)

type SSMClient struct {
	SSM  *ssm.AWSClients
	conf config.Config
}

func NewSSMClient(conf config.Config) (*SSMClient, error) {
	c, err := ssm.NewSSMClient(conf)
	if err != nil {
		return nil, err
	}

	return &SSMClient{
		SSM:  c,
		conf: conf,
	}, nil
}

func (c *SSMClient) KubeProxyOpen() error {
	// TODO: use port from config
	commands := []string{"runuser -l ec2-user -c \"kubectl proxy --port=8089 >/dev/null 2>&1 &\""}
	instanceIDs := []string{c.conf.AWS.InstanceID}

	err := c.SSM.SendCommand(commands, instanceIDs)
	if err != nil {
		return err
	}
	return nil
}

func (c *SSMClient) KubeProxyClose() error {
	commands := []string{"pkill -9 -f \"kubectl proxy\""}
	instanceIDs := []string{c.conf.AWS.InstanceID}

	err := c.SSM.SendCommand(commands, instanceIDs)
	if err != nil {
		return err
	}
	return nil
}

func (c *SSMClient) ForwordKubePort() error {
	instanceID := c.conf.AWS.InstanceID
	localPort := c.conf.AWS.LocalPort
	remotePort := c.conf.AWS.RemotePort

	err := c.SSM.ForwordPort(instanceID, localPort, remotePort, c.conf)
	if err != nil {
		return err
	}
	return nil
}
