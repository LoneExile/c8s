package kube

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/api/core/v1"
)

func (ps *KubeService) Pods(namespace string) (*v1.PodList , error) {
	pods, err := ps.KubeClient.Clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	return pods, nil
}
