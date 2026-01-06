package pvc

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type pvcData struct {
	Name         string   `json:"name,omitempty"`
	Namespace    string   `json:"namespace,omitempty"`
	Status       string   `json:"status,omitempty"`
	Capacity     string   `json:"capacity,omitempty"`
	AccessMode   []string `json:"accessMode,omitempty"`
	StorageClass string   `json:"storageClass,omitempty"`
	Volume       string   `json:"volume,omitempty"`
}

func ListPVCInNS(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pvcs, err := clientset.CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing pvc in %s: %v", ns, err)), nil
	}
	var output []pvcData
	for _, pvc := range pvcs.Items {
		qty := pvc.Spec.Resources.Requests[v1.ResourceStorage]
		output = append(output, pvcData{
			Name: pvc.Name,
			Namespace: pvc.Namespace,
			Capacity: qty.String(),
			Status: string(pvc.Status.Phase),
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListPVC (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []pvcData
	for _, namespace := range namespaces.Items {
		pvcs, err := clientset.CoreV1().PersistentVolumeClaims(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing pvc in %s: %v", namespace.Name, err)), nil
		}
		for _, pvc := range pvcs.Items {
			output = append(output, pvcData{
				Name: pvc.Name,
				Namespace: pvc.Namespace,
				Status: string(pvc.Status.Phase),
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetPVC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pvc, err := clientset.CoreV1().PersistentVolumeClaims(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting pvc in %s/%s: %v", ns, name, err)), nil
	}
	var accMode []string
	for _, mode := range pvc.Spec.AccessModes {
		accMode = append(accMode, string(mode))
	}
	qty := pvc.Spec.Resources.Requests[v1.ResourceStorage]
	
	output := pvcData{
		Name: pvc.Name,
		Namespace: pvc.Namespace,
		Capacity: qty.String(),
		AccessMode: accMode,
		StorageClass: *pvc.Spec.StorageClassName,
		Volume: pvc.Spec.VolumeName,
		Status: string(pvc.Status.Phase),
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeletePVC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().PersistentVolumeClaims(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting pvc in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully pvc %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdatePVC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	size, err := request.RequireString("size")
	if err != nil {
		output := fmt.Sprintf("Provide size for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pvc, err := clientset.CoreV1().PersistentVolumeClaims(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting pvc in %s/%s: %v", ns, name, err)), nil
	}

	qty, err := resource.ParseQuantity(size)
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Invalid pvc size: %s", size)), nil
	}

	pvc.Spec.Resources.Requests[v1.ResourceStorage] = qty

	updatePVC, err :=  clientset.CoreV1().PersistentVolumeClaims(ns).Update(context.TODO(), pvc, metav1.UpdateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in updating pvc in %s/%s with size %s: %v", ns, name, size, err)), nil
	}
	output := fmt.Sprintf("Successfully pvc %s/%s updated with size %s", updatePVC.Namespace, updatePVC.Name, size)
	return mcp.NewToolResultText(string(output)), nil
}

func CreatePVC(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	size, err := request.RequireString("size")
	if err != nil {
		output := fmt.Sprintf("Provide size for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	accessModes := request.GetString("accessMode", "ReadWriteOnce")
	storageClass, err := request.RequireString("storageClass")
	if err != nil {
		output := fmt.Sprintf("Provide storageClass name for pvc")
		return mcp.NewToolResultText(string(output)), nil
	}
	var accMode []v1.PersistentVolumeAccessMode
	am := strings.Split(accessModes, ",")
	for _, mode := range am {
		accMode = append(accMode, v1.PersistentVolumeAccessMode(mode))
	}

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	
	pvc := &v1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
            Name: name,
			Namespace: ns,
		},
		Spec: v1.PersistentVolumeClaimSpec{
			AccessModes: accMode,
			Resources: v1.VolumeResourceRequirements{
				Requests: v1.ResourceList{
					v1.ResourceStorage: resource.MustParse(size),
				},
			},
			StorageClassName: &storageClass,
		},
	}
	createPVC, err := clientset.CoreV1().PersistentVolumeClaims(ns).Create(context.TODO(), pvc, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating pvc %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully pvc %s/%s is created", createPVC.Namespace, createPVC.Name)
	return mcp.NewToolResultText(string(output)), nil
}