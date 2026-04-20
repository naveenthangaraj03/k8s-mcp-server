package createresource

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/restmapper"
)

func CreateResourceWithJson(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	jsondata, err := request.RequireString("jsondata")
	if err != nil {
		output := fmt.Sprintf("Provide jsonData to create resource")
		return mcp.NewToolResultText(string(output)), nil
	}
	_, dynamicClient, discoverClient, _, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	var obj unstructured.Unstructured
	if err = json.Unmarshal([]byte(jsondata), &obj); err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in unmarshal the json data: %v", err)), nil
	}

	groupResources, _ := restmapper.GetAPIGroupResources(discoverClient)
	mapper := restmapper.NewDiscoveryRESTMapper(groupResources)

	gvk := obj.GroupVersionKind()
	mapping, _ := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)

	var dr dynamic.ResourceInterface
	if mapping.Scope.Name() == "root" {
		dr = dynamicClient.Resource(mapping.Resource)
	} else {
		ns := obj.GetNamespace()
		if ns == "" {
			ns = "default"
		}
		dr = dynamicClient.Resource(mapping.Resource).Namespace(ns)
	}
	_, err = dr.Create(ctx, &obj, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating resource with jsondata: %v", err)), nil
	}
	return mcp.NewToolResultText(fmt.Sprintf("Successfully created resource with jsondata")), nil
}
