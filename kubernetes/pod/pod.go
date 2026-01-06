package pod

import (
	"fmt"
	"context"
	"encoding/json"
	"io"
	"strings"
	"strconv"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type podData struct {
	Name      string            `json:"name,omitempty"`
	Namespace string            `json:"namespace,omitempty"`
	Status    string            `json:"status,omitempty"`
	Labels    map[string]string `json:"labels,omitempty"`
	ContainerName []string      `json:"containerNames,omitempty"`
}

func ListPodInNS(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pods, err := clientset.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing pods in %s: %v", ns, err)), nil
	}
	var output []podData
	for _, pod := range pods.Items {
		output = append(output, podData{
			Name: pod.Name,
			Namespace: pod.Namespace,
			Status: string(pod.Status.Phase),
			Labels: pod.Labels,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListPod (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	labels := request.GetString("label", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}

	var output []podData
	for _, namespace := range namespaces.Items {
		pods, err := clientset.CoreV1().Pods(namespace.Name).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels,
		})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing pod in %s: %v", namespace.Name, err)), nil
		}

		for _, pod := range pods.Items {
			output = append(output, podData{
				Name: pod.Name,
				Namespace: pod.Namespace,
				Status: string(pod.Status.Phase),
				Labels: pod.Labels,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetPod (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pod, err := clientset.CoreV1().Pods(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting pods in %s/%s: %v", ns, name, err)), nil
	}

	cName := make([]string, 0, len(pod.Spec.Containers))
	
	for _, container := range pod.Spec.Containers {
		cName = append(cName, container.Name)
	}

	output := podData{
		Name: pod.Name,
		Namespace: pod.Namespace,
		Status: string(pod.Status.Phase),
		Labels: pod.Labels,
		ContainerName: cName,
	}
	
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeletePod (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.CoreV1().Pods(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting pods in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Pod %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdatePod (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels, err := request.RequireString("label")
	if err != nil {
		output := fmt.Sprintf("Provide label for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	pod, err := clientset.CoreV1().Pods(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting pods in %s/%s: %v", ns, name, err)), nil
	}
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
	pod.Labels = m
	updatePod, err := clientset.CoreV1().Pods(ns).Update(context.TODO(), pod, metav1.UpdateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in updating pod %s/%s with label %s: %v", ns, name, labels, err)), nil
	}
	output := fmt.Sprintf("Successfully pod %s/%s updated with label %s", updatePod.Namespace, updatePod.Name, labels)
	return mcp.NewToolResultText(string(output)), nil
}

func CreatePod (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	name,err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	containerNames, err := request.RequireString("containerNames")
	if err != nil {
		output := fmt.Sprintf("Provide container name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	containerImages, err := request.RequireString("containerImages")
	if err != nil {
		output := fmt.Sprintf("Provide image for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	containerPorts := request.GetString("containerPorts", "http:8080")
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

	cNames := strings.Split(containerNames, ",")
	cImages := strings.Split(containerImages, ",")
	cPorts := strings.Split(containerPorts, ",")

	if len(cNames) != len(cImages) {
		return mcp.NewToolResultText("container name and images counts are not matched"), nil
	}

	var containers []v1.Container

	for i := range cImages {
		var ports []v1.ContainerPort
		if i < len(cPorts) && cPorts[i] != "" {
			portDefs := strings.Split(cPorts[i], "|")
	
			for _, pd := range portDefs {
				parts := strings.SplitN(strings.TrimSpace(pd), ":", 2)
				if len(parts) != 2 {
					continue
				}
	
				portNum, err := strconv.Atoi(parts[1])
				if err != nil {
					continue
				}
	
				ports = append(ports, v1.ContainerPort{
					Name:          strings.TrimSpace(parts[0]),
					ContainerPort: int32(portNum),
				})
			}
		}
	
		if len(ports) == 0 {
			ports = append(ports, v1.ContainerPort{
				ContainerPort: 8080,
			})
		}
		containers = append(containers, v1.Container{
			Name:  strings.TrimSpace(cNames[i]),
			Image: strings.TrimSpace(cImages[i]),
			Ports: ports,
		})
	}

	pod := &v1.Pod{
        ObjectMeta: metav1.ObjectMeta{
            Name: name,
            Namespace: ns,
        },
        Spec: v1.PodSpec{
            Containers: containers,
        },
	}
	createPod, err := clientset.CoreV1().Pods(ns).Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating pod %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully pod %s/%s is created", createPod.Namespace, createPod.Name)
	return mcp.NewToolResultText(string(output)), nil
}

func PodLog (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	name,err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	tailLine  := request.GetInt("tailLine", 100)
	containerName, err := request.RequireString("containerName")
	if err != nil {
		output := fmt.Sprintf("Provide container name for pod")
		return mcp.NewToolResultText(string(output)), nil
	}
	count := int64(tailLine)
    podLogOptions := v1.PodLogOptions{
        Container: containerName,
        TailLines: &count,
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	req := clientset.CoreV1().Pods(ns).GetLogs(name, &podLogOptions)
	podLog, err := req.Stream(context.TODO())
	if err != nil {
		output := fmt.Sprintf("Error in streaming the log for Pod %s/%s", ns, name)
		return mcp.NewToolResultText(string(output)), nil
	}
	defer podLog.Close()
	body, err := io.ReadAll(podLog)
	if err != nil {
		output := fmt.Sprintf("Error in reading the log for Pod %s/%s", ns, name)
		return mcp.NewToolResultText(string(output)), nil
	}
	return mcp.NewToolResultText(string(body)), nil
}