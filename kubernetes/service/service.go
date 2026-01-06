package service

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"strconv"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"github.com/mark3labs/mcp-go/mcp"
)

type serviceData struct {
	Name          string            `json:"name,omitempty"`
	Namespace     string            `json:"namespace,omitempty"`
	Type          string            `json:"type,omitempty"`
	InternalIP    string            `json:"internalIP,omitempty"`
	ExternalIP    string            `json:"externalIP,omitempty"`
	SelectorLabel map[string]string `json:"selectorLabel,omitempty"`
}

func ListServiceInNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	services, err := clientset.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing service in %s: %v", ns, err)), nil
	}
	var output []serviceData
	for _, service := range services.Items {
		output = append(output, serviceData{
			Name: service.Name,
			Namespace: service.Namespace,
			Type: string(service.Spec.Type),
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListService (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output string
	for _, namespace := range namespaces.Items {
		services, err := clientset.CoreV1().Services(namespace.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing service in %s: %v", namespace.Name, err)), nil
		}
		var output []serviceData
		for _, service := range services.Items {
			output = append(output, serviceData{
				Name: service.Name,
				Namespace: service.Namespace,
				Type: string(service.Spec.Type),
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetService (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	service, err := clientset.CoreV1().Services(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting service in %s: %v", ns, err)), nil
	}

	var externalIP string
	if service.Spec.Type == "LoadBalancer" {
		temp := service.Status.LoadBalancer.Ingress
		externalIP = temp[0].IP
	}
	
	output := serviceData{
		Name: service.Name,
		Namespace: service.Namespace,
		Type: string(service.Spec.Type),
		InternalIP: service.Spec.ClusterIP,
		ExternalIP: externalIP,
		SelectorLabel: service.Spec.Selector,
	}
	
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteService (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().Services(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting service in %s/%s: %v", ns, name, err)), nil
	}
	
	output := fmt.Sprintf("Service %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdateService (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	selectorLabel := request.GetString("selectorLabel", "")
	svctype := request.GetString("type", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	service, err := clientset.CoreV1().Services(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting service in %s: %v", ns, err)), nil
	}
	if selectorLabel != "" {
		m := make(map[string]string)
		label := strings.Split(selectorLabel, ",")
		for _, lab := range label {
			kv := strings.SplitN(lab, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				m[key] = value
			}
		}
		service.Spec.Selector = m
		updateService, err := clientset.CoreV1().Services(ns).Update(context.TODO(), service, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating service in %s/%s: %v", ns, name, err)), nil
		}
		output := fmt.Sprintf("Successfully service %s/%s updated with label %s", updateService.Namespace, updateService.Name, selectorLabel)
		return mcp.NewToolResultText(string(output)), nil
	}
	if svctype != "" {
		service.Spec.Type = v1.ServiceType(svctype)
		updateService, err := clientset.CoreV1().Services(ns).Update(context.TODO(), service, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating service in %s/%s: %v", ns, name, err)), nil
		}
		output := fmt.Sprintf("Successfully service %s/%s updated with type %s", updateService.Namespace, updateService.Name, svctype)
		return mcp.NewToolResultText(string(output)), nil
	}
	output := fmt.Sprintf("Mentioned update in service %s/%s is not possible, we are supporting type and selectorLabelling", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateService (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	name,err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels, err := request.RequireString("selectorLabel")
	if err != nil {
		output := fmt.Sprintf("Provide selector label for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	svcType := request.GetString("svcType", "ClusterIP")
	svcPort, err := request.RequireString("svcPort")
	if err != nil {
		output := fmt.Sprintf("Provide svc port details for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	targetPort, err := request.RequireString("targetPort")
	if err != nil {
		output := fmt.Sprintf("Provide target port for service")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	lab := make(map[string]string)
	if labels != "" {
		deplabel := strings.Split(labels, ",")
		for _, label := range deplabel {
			kv := strings.SplitN(label, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				lab[key] = value
			}
		}
	}

	var ports []v1.ServicePort
	sPorts := strings.Split(svcPort, ",")
	tPorts := strings.Split(targetPort, ",")

	if len(sPorts) != len(tPorts) {
		return mcp.NewToolResultText("Service ports and target ports counts are not matched"), nil
	}

	for i := range sPorts {
		parts := strings.SplitN(strings.TrimSpace(sPorts[i]), ":", 2)
		portNum, err := strconv.Atoi(parts[1])
		if err != nil {
			continue
		}
		targetNum, err :=  strconv.Atoi(tPorts[i])
		if err != nil {
			continue
		}
		ports = append(ports, v1.ServicePort{
			Name:      strings.TrimSpace(parts[0]),
			Port:       int32(portNum),
			TargetPort: intstr.FromInt(targetNum),
		})
	}

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    lab,
		},
		Spec: v1.ServiceSpec{
			Selector: lab,
			Ports: ports,
			Type: v1.ServiceType(svcType),
		},
	}
	deployService, err := clientset.CoreV1().Services(ns).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating service in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully service %s/%s is created", deployService.Namespace, deployService.Name)
	return mcp.NewToolResultText(string(output)), nil
}