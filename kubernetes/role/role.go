package role

import (
	"fmt"
	"context"
	"encoding/json"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type roleData struct {
	Name         string    `json:"name,omitempty"`
	Namespace    string    `json:"namespace,omitempty"`
	Rules        []rules   ` json:"rules,omitempty"`
}

type rules struct {
	ApiGroups    []string  `json:"apiGroups,omitempty"`
	Resources    []string  `json:"resources,omitempty"`
	Verbs        []string  `json:"verbs,omitempty"`
}

func ListRoleInNS(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for role")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	roles, err := clientset.RbacV1().Roles(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing role in %s: %v", ns, err)), nil
	}
	var output []roleData
	for _, role := range roles.Items {
		output = append(output, roleData{
			Name: role.Name,
			Namespace: role.Namespace,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListRole(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []roleData
	for _, namespace := range namespaces.Items {
		roles, err := clientset.RbacV1().Roles(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing role in namespace %s: %v", namespace.Name, err)), nil
		}
		
		for _, role := range roles.Items {
			output = append(output, roleData{
				Name: role.Name,
				Namespace: role.Namespace,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetRole(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for role")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for role")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	role, err := clientset.RbacV1().Roles(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting role in %s/%s: %v", ns, name, err)), nil
	}

	var roleRules []rules

	for _, rule := range role.Rules {
		roleRules = append(roleRules, rules{
			ApiGroups: rule.APIGroups,
			Resources: rule.Resources,
			Verbs: rule.Verbs,
		})
	}
	
	output := roleData{
		Name: role.Name,
		Namespace: role.Namespace,
		Rules: roleRules,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}
