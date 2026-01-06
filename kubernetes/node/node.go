package node

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type nodeData struct {
	Name              string `json:"name,omitempty"`
	Status            string `json:"status,omitempty"`
	KubernetesVersion string `json:"kubernetesVersion,omitempty"`
	OS                string `json:"os,omitempty"`
	KernelVersion     string `json:"kernelVersion,omitempty"`
	Architecture      string `json:"architecture,omitempty"`
}

func ListNode (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing node: %v", err)), nil
	}
	var output []nodeData
	for _, node := range nodes.Items {
		var nodeStatus string
		for _, v := range node.Status.Conditions {
			if v.Type == "Ready" {
				if v.Status == "True"{
					nodeStatus = "Ready"
				} else {
					nodeStatus = "NotReady"
				}
			}
		}
		output = append(output, nodeData{
			Name: node.Name,
			Status: nodeStatus,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetNode (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for node")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting node: %v", err)), nil
	}
	var nodeStatus string
	for _, v := range node.Status.Conditions {
		if v.Type == "Ready" {
			if v.Status == "True"{
				nodeStatus = "Ready"
			} else {
				nodeStatus = "NotReady"
			}
		}
	}
	output := nodeData{
		Name: node.Name,
		Status: nodeStatus,
		KubernetesVersion: node.Status.NodeInfo.KubeletVersion,
		OS: node.Status.NodeInfo.OSImage,
		KernelVersion: node.Status.NodeInfo.KernelVersion,
		Architecture: node.Status.NodeInfo.Architecture,
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteNode (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for node")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().Nodes().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting node: %v", err)), nil
	}
	output := fmt.Sprintf("Node %s is deleted", name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdateNode (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for node")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels, err := request.RequireString("label")
	if err != nil {
		output := fmt.Sprintf("Provide label for node")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	node, err := clientset.CoreV1().Nodes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting node: %v", err)), nil
	}
	m := make(map[string]string)
	label := strings.Split(labels, ",")
	for _, lab := range label {
		kv := strings.SplitN(lab, "=", 2)
		if len(kv) == 2 {
			key := strings.TrimSpace(kv[0])
			value := strings.TrimSpace(kv[1])
			m[key] = value
		}
	}
	node.Labels = m
	updateNode, err := clientset.CoreV1().Nodes().Update(context.TODO(), node, metav1.UpdateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in updating node %s with label %s: %v", name, labels, err)), nil
	}
	output := fmt.Sprintf("Successfully node %s updated with label %s", updateNode.Name, labels)
	return mcp.NewToolResultText(string(output)), nil
}