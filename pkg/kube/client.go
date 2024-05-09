package kube

import (
	"net/http"
	"net/url"

	conf "c8s/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type Client struct {
	Clientset *kubernetes.Clientset
}

func NewClient(conf conf.Config) (*Client, error) {
	configPath := conf.KubeConfig.ConfigPath
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	if conf.Mode != "local" {
		proxyURL, err := url.Parse(conf.KubeConfig.ProxyUrl)
		if err != nil {
			return nil, err
		}
		config.WrapTransport = func(rt http.RoundTripper) http.RoundTripper {
			return &http.Transport{
				Proxy: http.ProxyURL(proxyURL),
			}
		}
	}

	config.Insecure = conf.KubeConfig.InsecureSkipVerify

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Client{Clientset: clientset}, nil
}
