package rolebinding

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type rbData struct {
	Name         string      `json:"name,omitempty"`
	Namespace    string      `json:"namespace,omitempty"`
	RoleRef      roleRef     `json:"roleRef,omitempty"`
	Subjects     []subjects  `json:"subjects,omitempty"`
}

type roleRef struct {
	ApiGroup     string  `json:"apiGroup,omitempty"`
	Kind         string  `json:"kind,omitempty"`
	Name         string  `json:"name,omitempty"`
}

type subjects struct {
	ApiGroup    string  `json:"apiGroup,omitempty"`
	Kind        string  `json:"kind,omitempty"`
	Name        string  `json:"name,omitempty"`
	Namespace   string  `json:"namespace,omitempty"`
}

func ListRBInNS(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for rolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	rbs, err := clientset.RbacV1().RoleBindings(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing rolebinding in %s: %v", ns, err)), nil
	}
	var output []rbData
	for _, rb := range rbs.Items {
		output = append(output, rbData{
			Name: rb.Name,
			Namespace: rb.Namespace,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []rbData
	for _, namespace := range namespaces.Items {
		rbs, err := clientset.RbacV1().RoleBindings(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing rolebinding in namespace %s: %v", namespace.Name, err)), nil
		}
		
		for _, rb := range rbs.Items {
			output = append(output, rbData{
				Name: rb.Name,
				Namespace: rb.Namespace,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for rolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for rolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	rb, err := clientset.RbacV1().RoleBindings(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting rolebinding in %s/%s: %v", ns, name, err)), nil
	}

	var saDetails []subjects

	for _, rolebind := range rb.Subjects {
		saDetails = append(saDetails, subjects{
			ApiGroup: rolebind.APIGroup,
			Kind: rolebind.Kind,
			Name: rolebind.Name,
			Namespace: rolebind.Namespace,
		})
	}

	var rRef roleRef
	rRef = roleRef{
		ApiGroup: rb.RoleRef.APIGroup,
		Kind: rb.RoleRef.Kind,
		Name: rb.RoleRef.Name,
	}
	
	output := rbData{
		Name: rb.Name,
		Namespace: rb.Namespace,
		RoleRef: rRef,
	    Subjects: saDetails,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}
