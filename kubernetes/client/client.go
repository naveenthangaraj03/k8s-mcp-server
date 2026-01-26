package client

import (
	"flag"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var kubeconfigPath string

func init() {
	flag.StringVar(&kubeconfigPath, "kubeconfigPath", "/root/.kube/conf", "Path to kubeconfig file")
}

func InitializeClients() (*kubernetes.Clientset, dynamic.Interface, error) {

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		return nil, nil, err
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return nil, nil, err
	}
	return clientset, dynamicClient, nil
}
