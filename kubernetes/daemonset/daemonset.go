package daemonset

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"strconv"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	"github.com/mark3labs/mcp-go/mcp"
)

type daemonsetData struct {
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	AvailableInstance string            `json:"availabeInstance,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	ContainerName     []string          `json:"containerName,omitempty"`
	ContainerImage    []string          `json:"containerImage,omitempty"`
}
func ListDaemonsetInNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	daemonsets, err := clientset.AppsV1().DaemonSets(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,
	})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing daemonsets in %s: %v", ns, err)), nil
	}
	var output []daemonsetData
	for _, daemonset := range daemonsets.Items {
		output = append(output, daemonsetData{
			Name: daemonset.Name,
			Namespace: daemonset.Namespace,
			AvailableInstance: fmt.Sprintf("%d/%d", daemonset.Status.NumberReady, daemonset.Status.UpdatedNumberScheduled),
			Labels: daemonset.Labels,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListDaemonset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []daemonsetData
	for _, namespace := range namespaces.Items {
		daemonsets, err := clientset.AppsV1().DaemonSets(namespace.Name).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels,
		})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing daemonsets in %s: %v", namespace.Name, err)), nil
		}
		for _, daemonset := range daemonsets.Items {
			output = append(output, daemonsetData{
				Name: daemonset.Name,
				Namespace: daemonset.Namespace,
				AvailableInstance: fmt.Sprintf("%d/%d", daemonset.Status.NumberReady, daemonset.Status.UpdatedNumberScheduled),
				Labels: daemonset.Labels,
			})
		}
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetDaemonset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	daemonset, err := clientset.AppsV1().DaemonSets(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting daemonsets in %s/%s: %v", ns, name, err)), nil
	}
	
	var cName []string
	var cImage []string
	for _, c := range daemonset.Spec.Template.Spec.Containers {
		cName = append(cName, c.Name)
		cImage = append(cImage, c.Image)
	}

	output := daemonsetData{
		Name: daemonset.Name,
		Namespace: daemonset.Namespace,
		AvailableInstance: fmt.Sprintf("%d/%d", daemonset.Status.NumberReady, daemonset.Status.UpdatedNumberScheduled),
		Labels: daemonset.Labels,
		ContainerName: cName,
		ContainerImage: cImage,
	}
	
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteDaemonset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.AppsV1().DaemonSets(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting daemonsets in %s: %v", ns, err)), nil
	}
	output := fmt.Sprintf("Daemonset %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdateDaemonset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	annotation := request.GetString("annotation", "")
	image := request.GetString("image", "")
	containerName := request.GetString("containerName", "")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	daemonset, err := clientset.AppsV1().DaemonSets(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting daemonsets in %s/%s: %v", ns, name, err)), nil
	}
	if labels != "" {
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
		daemonset.Labels = m
		updateDaemonset, err := clientset.AppsV1().DaemonSets(ns).Update(context.TODO(), daemonset, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating daemonset %s/%s with label %s: %v", ns, name, labels, err)), nil
		}
		output := fmt.Sprintf("Successfully daemonset %s/%s updated with label %s", updateDaemonset.Namespace, updateDaemonset.Name, labels)
		return mcp.NewToolResultText(string(output)), nil
	}
	if annotation != "" {
		m := make(map[string]string)
		annotations := strings.Split(annotation, ",")
		for _, ann := range annotations {
			kv := strings.SplitN(ann, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				m[key] = value
			}
		}
		daemonset.Annotations = m
		updateDaemonset, err := clientset.AppsV1().DaemonSets(ns).Update(context.TODO(), daemonset, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating daemonset %s/%s with annotation %s: %v", ns, name, annotation, err)), nil
		}
		output := fmt.Sprintf("Successfully daemonset %s/%s updated with annotaion %s", updateDaemonset.Namespace, updateDaemonset.Name, annotation)
		return mcp.NewToolResultText(string(output)), nil
	}
	if image != "" {
		if len(daemonset.Spec.Template.Spec.Containers) == 1 {
			daemonset.Spec.Template.Spec.Containers[0].Image = image
			updateDaemonset, err := clientset.AppsV1().DaemonSets(ns).Update(context.TODO(), daemonset, metav1.UpdateOptions{})
			if err != nil {
				return mcp.NewToolResultText(fmt.Sprintf("Error in updating daemonset %s/%s with image %s: %v", ns, name, image,  err)), nil
			}
			output := fmt.Sprintf("Successfully daemonset %s/%s updated with image %s", updateDaemonset.Namespace, updateDaemonset.Name, image)
			return mcp.NewToolResultText(string(output)), nil
		} else {
			if containerName == "" {
				output := fmt.Sprintf("Daemonset %s/%s has one than one container please provide the container name to update the image", ns, name)
				return mcp.NewToolResultText(string(output)), nil
			} else {
				var index int = -1
				for i, c := range daemonset.Spec.Template.Spec.Containers {
					if c.Name == containerName{
						index = i 
						break
					}
				}
				if index == -1 {
					output := fmt.Sprintf("Container name %s is not found in daemonset %s/%s ",containerName, ns, name)
					return mcp.NewToolResultText(string(output)), nil
				} else {
					daemonset.Spec.Template.Spec.Containers[index].Image = image
					updateDaemonset, err := clientset.AppsV1().DaemonSets(ns).Update(context.TODO(), daemonset, metav1.UpdateOptions{})
					if err != nil {
						return mcp.NewToolResultText(fmt.Sprintf("Error in updating daemonset %s/%s with image %s: %v", ns, name, image,  err)), nil
					}
					output := fmt.Sprintf("Successfully daemonset %s/%s updated with image %s", updateDaemonset.Namespace, updateDaemonset.Name, image)
					return mcp.NewToolResultText(string(output)), nil
				}
			}
		}
	}
	output := fmt.Sprintf("Mentioned update in daemonset %s/%s is not possible, we are supporting labelling, annotating and image", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateDaemonset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name,err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	containerNames, err := request.RequireString("containerNames")
	if err != nil {
		output := fmt.Sprintf("Provide container name for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	containerImages, err := request.RequireString("containerImages")
	if err != nil {
		output := fmt.Sprintf("Provide image for daemonset")
		return mcp.NewToolResultText(string(output)), nil
	}
	containerPorts := request.GetString("containerPorts", "http:8080")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	lab := make(map[string]string)
	if labels != "" {
		daemonsetlabel := strings.Split(labels, ",")
		for _, label := range daemonsetlabel {
			kv := strings.SplitN(label, "=", 2)
			if len(kv) == 2 {
				key := strings.TrimSpace(kv[0])
				value := strings.TrimSpace(kv[1])
				lab[key] = value
			}
		}
	}

	if len(lab) == 0 {
		lab["app"] = name
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

	daemonset := &appsv1.DaemonSet{
        ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: ns,
        },
        Spec: appsv1.DaemonSetSpec{
            Selector: &metav1.LabelSelector{
                MatchLabels: lab,
            },
            Template: v1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: lab,
                },
                Spec: v1.PodSpec{
					Containers: containers,
                },
            },
        },
	}
	deployDaemonset, err := clientset.AppsV1().DaemonSets(ns).Create(context.TODO(), daemonset, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deploying daemonset %s/%s: %v", ns, name,  err)), nil
	}
	output := fmt.Sprintf("Successfully daemonset %s/%s is created", deployDaemonset.Namespace, deployDaemonset.Name)
	return mcp.NewToolResultText(string(output)), nil
}