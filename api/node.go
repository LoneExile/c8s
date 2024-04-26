package api

import (
	ks "c8s/internal/service/kube"
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

type NodesOutput struct {
	Body struct {
		Nodes []string `json:"nodes" example:"[\"node1\", \"node2\"]" doc:"List of node names"`
	}
}

func Node(api huma.API, nodeService *ks.KubeService) {
	huma.Get(api, "/nodes", func(ctx context.Context, input *EmptyInput) (*NodesOutput, error) {
		nodes, err := nodeService.Nodes(metav1.NamespaceAll)
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Failed to list nodes")
		}

		nodeNames := make([]string, len(nodes.Items))
		for i, node := range nodes.Items {
			nodeNames[i] = node.Name
		}
		resp := &NodesOutput{}
		resp.Body.Nodes = nodeNames
		return resp, nil
	})
}
