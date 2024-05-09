package handlers

import (
	ks "c8s/internal/service/kube"
	"net/http"

	templates "c8s/internal/app/src"
	"c8s/internal/app/src/components"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Client struct {
	kubeService *ks.KubeService
}

func NewPodList(kubeService *ks.KubeService) *Client {
	p := &Client{
		kubeService: kubeService,
	}
	return p
}

func (h *Client) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	pods, err := h.kubeService.Pods(metav1.NamespaceAll)
	if err != nil {
		http.Error(w, "Failed to list pods", http.StatusInternalServerError)
		return
	}
	c := components.PodList(pods)
	err = templates.Layout(c, "Pods").Render(r.Context(), w)
	// err = components.PodList(pods).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (h *Client) Component(w http.ResponseWriter, r *http.Request) {

	pods, err := h.kubeService.Pods(metav1.NamespaceAll)
	if err != nil {
		http.Error(w, "Failed to list pods", http.StatusInternalServerError)
		return
	}
	err = components.PodList(pods).Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
