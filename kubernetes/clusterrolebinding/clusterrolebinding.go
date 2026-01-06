package clusterrolebinding

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type crbData struct {
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

func ListCRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	crbs, err := clientset.RbacV1().ClusterRoleBindings().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing clusterrolebinding: %v", err)), nil
	}
	var output []crbData
	for _, crb := range crbs.Items {
		output = append(output, crbData{
			Name: crb.Name,
			Namespace: crb.Namespace,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetCRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for clusterrolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	crb, err := clientset.RbacV1().ClusterRoleBindings().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting clusterrolebinding in %s: %v", name, err)), nil
	}

	var saDetails []subjects

	for _, crbind := range crb.Subjects {
		saDetails = append(saDetails, subjects{
			ApiGroup: crbind.APIGroup,
			Kind: crbind.Kind,
			Name: crbind.Name,
			Namespace: crbind.Namespace,
		})
	}

	var crRef roleRef
	crRef = roleRef{
		ApiGroup: crb.RoleRef.APIGroup,
		Kind: crb.RoleRef.Kind,
		Name: crb.RoleRef.Name,
	}
	
	output := crbData{
		Name: crb.Name,
		Namespace: crb.Namespace,
		RoleRef: crRef,
	    Subjects: saDetails,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}
