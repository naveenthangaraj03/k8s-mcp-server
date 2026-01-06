package configmap

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type cmData struct {
	Name      string            `json:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty"` 
	Data      map[string]string `json:"data,omitempty"`
}

func ListConfigmapInNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for configmap")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	configmaps, err := clientset.CoreV1().ConfigMaps(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing configmaps in %s: %v", ns, err)), nil
	}
	var output []cmData
	for _, configmap := range configmaps.Items {
		output = append(output, cmData{
			Name: configmap.Name,
			Namespace: configmap.Namespace,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListConfigmap (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []cmData
	for _, namespace := range namespaces.Items {
		configmaps, err := clientset.CoreV1().ConfigMaps(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing configmaps in %s: %v", namespace.Name, err)), nil
		}
		for _, configmap := range configmaps.Items {
			output = append(output, cmData{
				Name: configmap.Name,
				Namespace: configmap.Namespace,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetConfigmap (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for configmap")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for configmap")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	configmap, err := clientset.CoreV1().ConfigMaps(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting configmaps in %s/%s: %v", ns, name, err)), nil
	}
	output := cmData{
		Name: configmap.Name,
		Namespace: configmap.Namespace,
		Data: configmap.Data,
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteConfigmap (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for configmap")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for configmap")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().ConfigMaps(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting configmaps in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Configmap %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateConfigmap (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for configmap creation")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for configmap creation")
		return mcp.NewToolResultText(string(output)), nil
	}
	data, err := request.RequireString("data")
	if err != nil {
		output := fmt.Sprintf("Provide datas for configmap creation like password=kubernetes123,username=kubernetes")
		return mcp.NewToolResultText(string(output)), nil
	}

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	configmapData := make(map[string]string)
	
	cmData := strings.Split(data, ",")
	for _, datas := range cmData {
		kv := strings.SplitN(datas, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			configmapData[key] = value
		}
	}

	configmap := &v1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
		},
		Data: configmapData,
	}
	createConfigmap, err := clientset.CoreV1().ConfigMaps(ns).Create(context.TODO(), configmap, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating configmap in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully configmap %s/%s is created", createConfigmap.Namespace, createConfigmap.Name)
	return mcp.NewToolResultText(string(output)), nil
}