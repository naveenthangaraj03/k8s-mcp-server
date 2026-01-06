package namespace

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

type namespaceData struct {
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
}

func ListNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []namespaceData
	for _, namespace := range namespaces.Items {
		output = append(output, namespaceData{
			Name: namespace.Name,
			Status: string(namespace.Status.Phase),
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide namespace name to get")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in gettting the namespace %s: %v", name, err)), nil
	}
	output := namespaceData{
		Name: namespace.Name,
		Status: string(namespace.Status.Phase),
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide namespace name to delete")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting the namespace %s: %v", name, err)), nil
	}
	output := fmt.Sprintf("Namespace %s is deleted", name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdateNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide namespace name to update")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	annotation := request.GetString("annotation", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in gettting the namespace %s: %v", name, err)), nil
	}
	if labels != "" {
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
		namespace.Labels = m
		updateNamespace, err := clientset.CoreV1().Namespaces().Update(context.TODO(), namespace, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating namesapce %s with label %s: %v", name, labels, err)), nil
		}
		output := fmt.Sprintf("Successfully namespace %s updated with label %s", updateNamespace.Name, labels)
		return mcp.NewToolResultText(string(output)), nil
	}
	if annotation != "" {
		m := make(map[string]string)
		annotations := strings.Split(annotation, ",")
		for _, ann := range annotations {
			kv := strings.SplitN(ann, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				m[key] = value
			}
		}
		namespace.Annotations = m
		updateNamespace, err := clientset.CoreV1().Namespaces().Update(context.TODO(), namespace, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating namespace %s with annotation %s: %v", name, annotation, err)), nil
		}
		output := fmt.Sprintf("Successfully namespace %s updated with annotaion %s",  updateNamespace.Name, annotation)
		return mcp.NewToolResultText(string(output)), nil
	}
	output := fmt.Sprintf("Mentioned update in namespace %s is not possible, we are supporting labelling and  annotating",  name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide namespace name to create")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	lab := make(map[string]string)
	if labels != "" {
		nslabel := strings.Split(labels, ",")
		for _, label := range nslabel {
			kv := strings.SplitN(label, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				lab[key] = value
			}
		}
	}

	namespace := &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		    Labels: lab,
		},
	}

	createNamespace, err := clientset.CoreV1().Namespaces().Create(context.TODO(), namespace, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating namespace %s: %v", name, err)), nil
	}
	output := fmt.Sprintf("Successfully namespace %s is created",  createNamespace.Name)
	return mcp.NewToolResultText(string(output)), nil
}