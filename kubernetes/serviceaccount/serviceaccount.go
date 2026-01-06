package serviceaccount

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type saData struct {
	Name      string            `json:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
}

func ListSAInNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	sAccount, err := clientset.CoreV1().ServiceAccounts(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing service accounts in %s: %v", ns, err)), nil
	}
	var output []saData
	for _, sa := range sAccount.Items {
		output = append(output, saData{
			Name: sa.Name,
			Namespace: sa.Namespace,
			Labels: sa.Labels,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListSA (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []saData
	for _, namespace := range namespaces.Items {
		sAccount, err := clientset.CoreV1().ServiceAccounts(namespace.Name).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels,
		})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing service accounts in %s: %v", namespace.Name, err)), nil
		}
		
		for _, sa := range sAccount.Items {
			output = append(output, saData{
				Name: sa.Name,
				Namespace: sa.Namespace,
				Labels: sa.Labels,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetSA (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	sAccount, err := clientset.CoreV1().ServiceAccounts(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting service accounts in %s/%s: %v", ns, name, err)), nil
	}
	
	output := saData{
		Name: sAccount.Name,
		Namespace: sAccount.Namespace,
		Labels: sAccount.Labels,
	}


	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteSA (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	err = clientset.CoreV1().ServiceAccounts(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting service accounts in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("ServiceAccount %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateSA (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name to create Service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace to create Service account")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	lab := make(map[string]string)
	if labels != "" {
		salabel := strings.Split(labels, ",")
		for _, label := range salabel {
			kv := strings.SplitN(label, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				lab[key] = value
			}
		}
	}

	serviceaccount := &v1.ServiceAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: ns,
		    Labels: lab,
		},
	}

	createServiceAccount, err := clientset.CoreV1().ServiceAccounts(ns).Create(context.TODO(), serviceaccount, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating service account %s/%s: %v", ns , name, err)), nil
	}
	output := fmt.Sprintf("Successfully serviceAccount %s/%s is created", createServiceAccount.Namespace, createServiceAccount.Name)
	return mcp.NewToolResultText(string(output)), nil
}