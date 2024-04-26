package kube

import (
	"c8s/pkg/kube"
)

type KubeService struct {
	KubeClient *kube.Client
}

func NewService(client *kube.Client) *KubeService {
	return &KubeService{KubeClient: client}
}
