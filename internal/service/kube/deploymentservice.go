package kube

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/api/core/v1"
)

func (ps *KubeService) Nodes(namespace string) (*v1.NodeList , error) {
	nodes, err := ps.KubeClient.Clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return nodes, nil
}
