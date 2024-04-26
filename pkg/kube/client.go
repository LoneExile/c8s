package kube

import (
	"flag"
	"path/filepath"

	conf "c8s/config"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func InitKube(conf conf.Config) *kubernetes.Clientset {
	var kubeconfig *string
	home := homedir.HomeDir()
	// kubeconfig = flag.String("kubeconfig", filepath.Join(".kube", "tt"), "")
	kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	config.Insecure = conf.KubeConfig.InsecureSkipVerify

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return clientset
}

type Client struct {
	Clientset *kubernetes.Clientset
}

func NewClient(conf conf.Config) (*Client, error) {
	configPath := filepath.Join(homedir.HomeDir(), ".kube", "config")
	config, err := clientcmd.BuildConfigFromFlags("", configPath)
	if err != nil {
		return nil, err
	}

	config.Insecure = conf.KubeConfig.InsecureSkipVerify

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &Client{Clientset: clientset}, nil
}
