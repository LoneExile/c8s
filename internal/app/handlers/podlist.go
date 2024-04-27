package handlers

import (
	templates "c8s/internal/app/src"
	ks "c8s/internal/service/kube"
	"net/http"

	"c8s/internal/app/src/components"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodList struct {
	kubeService *ks.KubeService
}

func NewPodList(kubeService *ks.KubeService) *PodList {
	p := &PodList{
		kubeService: kubeService,
	}
	return p
}

func (h *PodList) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pods, err := h.kubeService.Pods(metav1.NamespaceAll)
	if err != nil {
		http.Error(w, "Failed to list pods", http.StatusInternalServerError)
		return
	}
	c := components.PodList(pods)
	err = templates.Layout(c, "Pods").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
