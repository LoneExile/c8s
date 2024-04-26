package api

import (
	ks "c8s/internal/service/kube"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type EmptyInput struct{}

type PodsOutput struct {
	Body struct {
		Pods []string `json:"pods" example:"[\"pod1\", \"pod2\"]" doc:"List of pod names"`
	}
}

func Pod(api huma.API, podService *ks.KubeService) {
	huma.Get(api, "/pods", func(ctx context.Context, input *EmptyInput) (*PodsOutput, error) {
		pods, err := podService.Pods(metav1.NamespaceAll)
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Failed to list pods")
		}

		podNames := make([]string, len(pods.Items))
		for i, pod := range pods.Items {
			podNames[i] = pod.Name
		}
		resp := &PodsOutput{}
		resp.Body.Pods = podNames
		return resp, nil
	})
}
