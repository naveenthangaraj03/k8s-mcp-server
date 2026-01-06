package pv

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type pvData struct {
	Name         string   `json:"name,omitempty"`
	Namespace    string   `json:"namespace,omitempty"`
	Status       string   `json:"status,omitempty"`
	Capacity     string   `json:"capacity,omitempty"`
	AccessMode   []string `json:"accessMode,omitempty"`
	StorageClass string   `json:"storageClass,omitempty"`
}

func ListPV(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pvolume, err := clientset.CoreV1().PersistentVolumes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing pv: %v", err)), nil
	}
	var output []pvData
	for _, pv := range pvolume.Items {
		qty := pv.Spec.Capacity[v1.ResourceStorage]
		output = append(output, pvData{
			Name: pv.Name,
			Namespace: pv.Namespace,
			Capacity: qty.String(),
			Status: string(pv.Status.Phase),
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetPV(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pv")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pv, err := clientset.CoreV1().PersistentVolumes().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting pv in %s: %v", name, err)), nil
	}

	var accMode []string
	for _, mode := range pv.Spec.AccessModes {
		accMode = append(accMode, string(mode))
	}
	qty := pv.Spec.Capacity[v1.ResourceStorage]
	
	output := pvData{
		Name: pv.Name,
		Namespace: pv.Namespace,
		Capacity: qty.String(),
		AccessMode: accMode,
		StorageClass: pv.Spec.StorageClassName,
		Status: string(pv.Status.Phase),
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeletePV(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pv")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().PersistentVolumes().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting pv %s: %v", name, err)), nil
	}
	output := fmt.Sprintf("Successfully pv %s is deleted", name)
	return mcp.NewToolResultText(string(output)), nil
}