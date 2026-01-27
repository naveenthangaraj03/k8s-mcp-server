package clusterrolebinding

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

type crbData struct {
	Name     string     `json:"name,omitempty"`
	RoleRef  roleRef    `json:"roleRef,omitempty"`
	Subjects []subjects `json:"subjects,omitempty"`
}

type roleRef struct {
	ApiGroup string `json:"apiGroup,omitempty"`
	Kind     string `json:"kind,omitempty"`
	Name     string `json:"name,omitempty"`
}

type subjects struct {
	ApiGroup  string `json:"apiGroup,omitempty"`
	Kind      string `json:"kind,omitempty"`
	Name      string `json:"name,omitempty"`
	Namespace string `json:"namespace,omitempty"`
}

func ListCRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, _, _, err := client.InitializeClients()
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
	clientset, _, _, err := client.InitializeClients()
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
			ApiGroup:  crbind.APIGroup,
			Kind:      crbind.Kind,
			Name:      crbind.Name,
			Namespace: crbind.Namespace,
		})
	}

	var crRef roleRef
	crRef = roleRef{
		ApiGroup: crb.RoleRef.APIGroup,
		Kind:     crb.RoleRef.Kind,
		Name:     crb.RoleRef.Name,
	}

	output := crbData{
		Name:     crb.Name,
		RoleRef:  crRef,
		Subjects: saDetails,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteCRB(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for clusterrolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, _, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.RbacV1().ClusterRoleBindings().Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting clusterrolebinding named %s: %v", name, err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully deleted clusterrolebinding named %s", name)), nil
}

func CreateCRBWithJson(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	jsondata, err := request.RequireString("jsondata")
	if err != nil {
		output := fmt.Sprintf("Provide jsonData for clusterrolebinding")
		return mcp.NewToolResultText(string(output)), nil
	}
	_, dynamicClient, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	resourceId := schema.GroupVersionResource{
		Group:    "rbac.authorization.k8s.io",
		Version:  "v1",
		Resource: "clusterrolebindings",
	}

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(jsondata), &obj); err != nil {
		return nil, err
	}

	unstructuredObj := &unstructured.Unstructured{Object: obj}

	_, err = dynamicClient.Resource(resourceId).Create(ctx, unstructuredObj, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating clusterrolebinding with jsondata: %v", err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully created clusterrolebinding with jsondata")), nil
}
