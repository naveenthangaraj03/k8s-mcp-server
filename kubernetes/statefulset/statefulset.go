package statefulset

import (
	"fmt"
	"context"
	"encoding/json"
	"strings"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/client"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/api/core/v1"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/api/resource"
	"github.com/mark3labs/mcp-go/mcp"
)

type stsData struct{
	Name              string            `json:"name,omitempty"`
	Namespace         string            `json:"namespace,omitempty"`
	AvailableInstance string            `json:"availableInstance,omitempty"`
	Labels            map[string]string `json:"labels,omitempty"`
	ContainerName     []string          `json:"containerName,omitempty"`
	ContainerImage    []string          `json:"containerImage,omitempty"`
}

func ListStatefulsetInNS (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	statefulsets, err := clientset.AppsV1().StatefulSets(ns).List(context.TODO(), metav1.ListOptions{
		LabelSelector: labels,	
	})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing statefulset in %s: %v", ns, err)), nil
	}
	var output []stsData
	for _, statefulset := range statefulsets.Items {
		output = append(output, stsData{
			Name: statefulset.Name,
			Namespace: statefulset.Namespace,
			AvailableInstance: fmt.Sprintf(	"%d/%d", statefulset.Status.AvailableReplicas, *statefulset.Spec.Replicas,),
			Labels: statefulset.Labels,
		})
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func ListStatefulset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	labels := request.GetString("label", "")

	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in listing namespace: %v", err)), nil
	}
	var output []stsData
	for _, namespace := range namespaces.Items {
		statefulsets, err := clientset.AppsV1().StatefulSets(namespace.Name).List(context.TODO(), metav1.ListOptions{
			LabelSelector: labels,
		})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in listing statefulset in %s: %v", namespace.Name, err)), nil
		}
		for _, statefulset := range statefulsets.Items {
			output = append(output, stsData{
				Name: statefulset.Name,
				Namespace: statefulset.Namespace,
				AvailableInstance: fmt.Sprintf(	"%d/%d", statefulset.Status.AvailableReplicas, *statefulset.Spec.Replicas,),
				Labels: statefulset.Labels,
			})
		}	
	}
	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func GetStatefulset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide names for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	statefulset, err := clientset.AppsV1().StatefulSets(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting statefulset in %s: %v", ns, err)), nil
	}

	var cName []string
	var cImage []string
	for _, c := range statefulset.Spec.Template.Spec.Containers {
		cName = append(cName, c.Name)
		cImage = append(cImage, c.Image)
	}

	output := stsData{
		Name: statefulset.Name,
		Namespace: statefulset.Namespace,
		AvailableInstance: fmt.Sprintf(	"%d/%d", statefulset.Status.AvailableReplicas, *statefulset.Spec.Replicas,),
		Labels: statefulset.Labels,
		ContainerName: cName,
		ContainerImage: cImage,
	}

	mcpOutput, err := json.MarshalIndent(output, "", " ")
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in marshalling: %v", err)), nil
	}
	return mcp.NewToolResultText(string(mcpOutput)), nil
}

func DeleteStatefulset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	err = clientset.AppsV1().StatefulSets(ns).Delete(context.TODO(), name, metav1.DeleteOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in deleting statefulset in %s: %v", ns, err)), nil
	}

	output := fmt.Sprintf("Statefulset %s/%s is deleted", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func UpdateStatefulset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name, err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	annotation := request.GetString("annotation", "")
	image := request.GetString("image", "")
	containerName := request.GetString("containerName", "")
	replica := request.GetInt("replica", -1)
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}
	statefulset, err := clientset.AppsV1().StatefulSets(ns).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in getting statefulset in %s: %v", ns, err)), nil
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
		statefulset.Labels = m
		updateStatefulset, err := clientset.AppsV1().StatefulSets(ns).Update(context.TODO(), statefulset, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating statefulset %s/%s with label %s: %v", ns, name, labels, err)), nil
		}
		output := fmt.Sprintf("Successfully statefulset %s/%s updated with label %s", updateStatefulset.Namespace, updateStatefulset.Name, labels)
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
		statefulset.Annotations = m
		updateStatefulset, err := clientset.AppsV1().StatefulSets(ns).Update(context.TODO(), statefulset, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating statefulset  %s/%s with annotation %s: %v", ns, name, annotation, err)), nil
		}
		output := fmt.Sprintf("Successfully statefulset %s/%s updated with annotaion %s", updateStatefulset.Namespace, updateStatefulset.Name, annotation)
		return mcp.NewToolResultText(string(output)), nil
	}
	if image != "" {
		if len(statefulset.Spec.Template.Spec.Containers) == 1 {
			statefulset.Spec.Template.Spec.Containers[0].Image = image
			updateStatefulset, err := clientset.AppsV1().StatefulSets(ns).Update(context.TODO(), statefulset, metav1.UpdateOptions{})
			if err != nil {
				return mcp.NewToolResultText(fmt.Sprintf("Error in updating statefulset %s/%s with image %s: %v", ns, name, image,  err)), nil
			}
			output := fmt.Sprintf("Successfully statefulset %s/%s updated with image %s", updateStatefulset.Namespace, updateStatefulset.Name, image)
			return mcp.NewToolResultText(string(output)), nil
		} else {
			if containerName == "" {
				output := fmt.Sprintf("Statefulset %s/%s has one than one container please provide the container name to update the image", ns, name)
				return mcp.NewToolResultText(string(output)), nil
			} else {
				var index int = -1
				for i, c := range statefulset.Spec.Template.Spec.Containers {
					if c.Name == containerName{
						index = i 
						break
					}
				}
				if index == -1 {
					output := fmt.Sprintf("Container name %s is not found in statefulset %s/%s ",containerName, ns, name)
					return mcp.NewToolResultText(string(output)), nil
				} else {
					statefulset.Spec.Template.Spec.Containers[index].Image = image
					updateStatefulset, err := clientset.AppsV1().StatefulSets(ns).Update(context.TODO(), statefulset, metav1.UpdateOptions{})
					if err != nil {
						return mcp.NewToolResultText(fmt.Sprintf("Error in updating statefulset %s/%s with image %s: %v", ns, name, image, err)), nil
					}
					output := fmt.Sprintf("Successfully statefulset %s/%s updated with image %s", updateStatefulset.Namespace, updateStatefulset.Name, image)
					return mcp.NewToolResultText(string(output)), nil
				}
			}
		}
	}
	if replica > -1 {
		replicas := int32(replica)
		statefulset.Spec.Replicas = &replicas
		updateStatefulset, err := clientset.AppsV1().StatefulSets(ns).Update(context.TODO(), statefulset, metav1.UpdateOptions{})
		if err != nil {
			return mcp.NewToolResultText(fmt.Sprintf("Error in updating statefulset %s/%s with replica %d: %v", ns, name, replica, err)), nil
		}
		output := fmt.Sprintf("Successfully statefulset %s/%s updated with replica %d", updateStatefulset.Namespace, updateStatefulset.Name, replica)
		return mcp.NewToolResultText(string(output)), nil
	}
	output := fmt.Sprintf("Mentioned update in statefulset %s/%s is not possible, we are supporting labelling, annotating, replica and image", ns, name)
	return mcp.NewToolResultText(string(output)), nil
}

func CreateStatefulset (ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	ns, err := request.RequireString("namespace")
	if err != nil {
		output := fmt.Sprintf("Provide namespace for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	name,err := request.RequireString("name")
	if err != nil {
		output := fmt.Sprintf("Provide name for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	labels := request.GetString("label", "")
	replica := request.GetInt("replica", 1)
	containerName := request.GetString("containerNames", name)
	containerImage, err := request.RequireString("containerImages")
	if err != nil {
		output := fmt.Sprintf("Provide image for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	containerPort  := request.GetInt("containerPorts", 8080)
	storageValue, err := request.RequireString("storageValue")
	if err != nil {
		output := fmt.Sprintf("Provide storage value for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	mountPath, err := request.RequireString("mountPath")
	if err != nil {
		output := fmt.Sprintf("Provide mount path for statefulset")
		return mcp.NewToolResultText(string(output)), nil
	}
	pvcName := request.GetString("pvcName", name)
	svcPort  := request.GetInt("svcPort", 8080)
	svcType := request.GetString("svcType", "ClusterIP")
	clientset, err := client.InitializeClients()
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in intialize client: %v", err)), nil
	}

	var dsReplica int32
	dsReplica = int32(replica)
	lab := make(map[string]string)
	if labels != "" {
		dslabel := strings.Split(labels, ",")
		for _, label := range dslabel {
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

	service := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: ns,
			Labels:    lab,
		},
		Spec: v1.ServiceSpec{
			Selector: lab,
			Ports: []v1.ServicePort{
				{
					Name:       name,
					Port:       int32(svcPort),
					TargetPort: intstr.FromInt(containerPort),
				},
			},
			Type: v1.ServiceType(svcType),
		},
	}

	deployService, err := clientset.CoreV1().Services(ns).Create(context.TODO(), service, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating service for sts in %s/%s: %v", ns, name, err)), nil
	}

	statefulset := &appsv1.StatefulSet{
        ObjectMeta: metav1.ObjectMeta{
			Name: name,
			Namespace: ns,
        },
        Spec: appsv1.StatefulSetSpec{
            ServiceName: name,
            Replicas:    &dsReplica,
            Selector: &metav1.LabelSelector{
                MatchLabels: lab,
            },
            Template: v1.PodTemplateSpec{
                ObjectMeta: metav1.ObjectMeta{
                    Labels: lab,
                },
                Spec: v1.PodSpec{
                    Containers: []v1.Container{
                        {
                            Name:  containerName,
                            Image: containerImage,
                            Ports: []v1.ContainerPort{
                                {
                                    ContainerPort: int32(containerPort),
                                    Name:          name,
                                },
                            },
                            VolumeMounts: []v1.VolumeMount{
                                {
                                    Name:      pvcName,
                                    MountPath: mountPath,
                                },
                            },
                        },
                    },
                },
            },
            VolumeClaimTemplates: []v1.PersistentVolumeClaim{
                {
                    ObjectMeta: metav1.ObjectMeta{
                        Name: pvcName,
                    },
                    Spec: v1.PersistentVolumeClaimSpec{
                        AccessModes: []v1.PersistentVolumeAccessMode{
                            v1.ReadWriteOnce,
                        },
                        Resources: v1.VolumeResourceRequirements{
                            Requests: v1.ResourceList{
                                v1.ResourceStorage: resource.MustParse(storageValue),
                            },
                        },
                    },
                },
            },
        },
	}
	deployStatefulset, err := clientset.AppsV1().StatefulSets(ns).Create(context.TODO(), statefulset, metav1.CreateOptions{})
	if err != nil {
		return mcp.NewToolResultText(fmt.Sprintf("Error in creating statefulset in %s/%s: %v", ns, name, err)), nil
	}
	output := fmt.Sprintf("Successfully statefulset %s/%s is created with service %s", deployStatefulset.Namespace, deployStatefulset.Name, deployService.Name)
	return mcp.NewToolResultText(string(output)), nil
}