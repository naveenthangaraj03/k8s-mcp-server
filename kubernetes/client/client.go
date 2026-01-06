package client

import (
	"k8s.io/client-go/kubernetes"
    "k8s.io/client-go/tools/clientcmd"
    "flag"
)

var kubeconfigPath string

func init() {
    flag.StringVar(&kubeconfigPath, "kubeconfigPath", "/root/.kube/conf", "Path to kubeconfig file")
}

func InitializeClients() (*kubernetes.Clientset, error) {

    config, err := clientcmd.BuildConfigFromFlags("" , kubeconfigPath)
    if err != nil {
        return nil, err
    }
    clientset, err := kubernetes.NewForConfig(config)
    if err != nil {
        return nil, err
    }
    return clientset, nil
}
