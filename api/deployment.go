package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	ks "c8s/internal/service/kube"

	"github.com/danielgtaylor/huma/v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RestartDeploymentOutput struct {
	Body struct {
		Message string `json:"message"`
	}
}

type RestartDeploymentInput struct {
	Body struct {
		Namespace      string `json:"namespace" doc:"Namespace of the deployment"`
		DeploymentName string `json:"deploymentName" doc:"Name of the deployment"`
	}
}

func RestartDeployment(api huma.API, kubeService *ks.KubeService) {
	huma.Register(api, huma.Operation{
		OperationID:   "restartDeployment",
		Summary:       "Restart a deployment",
		Method:        http.MethodPost,
		Path:          "/api/deployments/restart",
		Tags:          []string{"deployments"},
		DefaultStatus: http.StatusOK,
	}, func(ctx context.Context, input *RestartDeploymentInput) (*RestartDeploymentOutput, error) {
		kkc := kubeService.KubeClient.Clientset

		fmt.Println("Namespace: ", input.Body.Namespace)
		fmt.Println("Deployment Name: ", input.Body.DeploymentName)

		deployment, err := kkc.AppsV1().Deployments(input.Body.Namespace).Get(ctx, input.Body.DeploymentName, metav1.GetOptions{})
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Failed to get deployment")
		}

		if deployment.Spec.Template.Annotations == nil {
			deployment.Spec.Template.Annotations = make(map[string]string)
		}
		deployment.Spec.Template.Annotations["kubectl.kubernetes.io/restartedAt"] = time.Now().Format(time.RFC3339)

		_, err = kkc.AppsV1().Deployments(input.Body.Namespace).Update(ctx, deployment, metav1.UpdateOptions{})
		if err != nil {
			return nil, huma.NewError(http.StatusInternalServerError, "Failed to restart deployment")
		}

		response := &RestartDeploymentOutput{}
		response.Body.Message = "Deployment restarted successfully"
		return response, nil
	})
}
