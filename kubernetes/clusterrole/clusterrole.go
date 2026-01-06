package clusterrole

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type crData struct {
	Name         string    `json:"name,omitempty"`
	Namespace    string    `json:"namespace,omitempty"`
	Rules        []rules   `json:"rules,omitempty"`
}

type rules struct {
	ApiGroups    []string  `json:"apiGroups,omitempty"`
	Resources    []string  `json:"resources,omitempty"`
	Verbs        []string  `json:"verbs,omitempty"`
}

func ListCR(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	crs, err := clientset.RbacV1().ClusterRoles().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing clusterrole: %v", err)), nil
	}
	var output []crData
	for _, cr := range crs.Items {
		output = append(output, crData{
			Name: cr.Name,
			Namespace: cr.Namespace,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetCR(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for clusterrole")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	cr, err := clientset.RbacV1().ClusterRoles().Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting clusterrole in %s: %v", name, err)), nil
	}

	var crRules []rules

	for _, rule := range cr.Rules {
		crRules = append(crRules, rules{
			ApiGroups: rule.APIGroups,
			Resources: rule.Resources,
			Verbs: rule.Verbs,
		})
	}
	
	output := crData{
		Name: cr.Name,
		Namespace: cr.Namespace,
		Rules: crRules,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}