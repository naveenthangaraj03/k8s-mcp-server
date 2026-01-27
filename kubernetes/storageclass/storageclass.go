package storageclass

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type scData struct {
	Name          string `json:"name,omitempty"`
	Provisioner   string `json:"provisioner,omitempty"`
	ReclaimPolicy string `json:"reclaimPolicy,omitempty"`
}

func ListSC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, _, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	sc, err := clientset.StorageV1().StorageClasses().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing storageclass: %v", err)), nil
	}
	var output []string
	for _, sclass := range sc.Items {
		output = append(output, sclass.Name)
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetSC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for storage class")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, _, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	sc, err := clientset.StorageV1().StorageClasses().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting storageclass %s: %v", name, err)), nil
	}
	output := scData{
		Name:          sc.Name,
		Provisioner:   sc.Provisioner,
		ReclaimPolicy: string(*sc.ReclaimPolicy),
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteSC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for storageclass")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, _, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.StorageV1().StorageClasses().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting storageclass named %s: %v", name, err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully deleted storageclass named %s", name)), nil
}

func CreateSCWithJson(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	jsondata, err := request.RequireString("jsondata")
	if err != nil {
		output := fmt.Sprintf("Provide jsonData for storageclass")
		return mcp.NewToolResultText(string(output)), nil
	}
	_, dynamicClient, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	resourceId := schema.GroupVersionResource{
		Group:    "storage.k8s.io",
		Version:  "v1",
		Resource: "storageclasses",
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(jsondata), &obj); err != nil {
		return nil, err
	}

	unstructuredObj := &unstructured.Unstructured{Object: obj}

	_, err = dynamicClient.Resource(resourceId).Create(ctx, unstructuredObj, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating storageclass with jsondata: %v", err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully created storageclass with jsondata")), nil
}
