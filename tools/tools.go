package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
)

var ListPodInNamespace = mcp.NewTool(
	"list-pod-in-namespace",
    mcp.WithDescription("List the pod in particular namespace with status, label and instance"),
    mcp.WithString(
		"namespace",
		mcp.Required(),
        mcp.Description("The namespace in which the pod should be listed"),
	),
	mcp.WithString(
		"label", 
		mcp.Description("Only return pods matching this label selector"),
	),
)

var ListPod = mcp.NewTool(
	"list-pod",
	mcp.WithDescription("List the pod in all namespaces with status, label and instance"),
	mcp.WithString(
		"label", 
		mcp.Description("Only return pods matching this label selector"),
	),
)

var GetPod = mcp.NewTool(
	"get-pod",
    mcp.WithDescription("Get the pod in particular namespace with status, label and instance"),
    mcp.WithString(
		"namespace",
		mcp.Required(),
        mcp.Description("The namespace in which the pod exists"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
        mcp.Description("The name of the pod to get details"),
	),
)

var DeletePod = mcp.NewTool(
	"delete-pod",
    mcp.WithDescription("Delete the pod in particular namespace"),
    mcp.WithString(
		"namespace",
		mcp.Required(),
        mcp.Description("The namespace in which the pod to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
        mcp.Description("The name of the pod to be deleted"),
	),
)

var UpdatePod = mcp.NewTool(
	"update-pod",
	mcp.WithDescription("Update the pod in particular namespace like label changes"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the pod to be updated"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pod to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Required(),
		mcp.Description("Label to be updated"),
	),
)

var CreatePod = mcp.NewTool(
	"create-pod",
	mcp.WithDescription("Create the pod in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the pod to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pod to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be added in that pod"),
	),		
	mcp.WithString(
		"containerNames",
		mcp.Required(),
		mcp.Description("Container Names for the pod"),
	),
	mcp.WithString(
		"containerImages",
		mcp.Required(),
		mcp.Description("Container Image for the pod"),
	),
	mcp.WithString(
		"containerPorts",
		mcp.Description("Container port details for the pod"),
	),
)

var PodLog = mcp.NewTool(
	"pod-log",
	mcp.WithDescription("Get the log for particular pod in the namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the pod is present"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pod to get log"),
	),
	mcp.WithNumber(
		"tailLine",
		mcp.Description("Number of log line to get"),
	),
	mcp.WithString(
		"containerName",
		mcp.Required(),
		mcp.Description("Container Names for the pod to get log"),
	),
)

var ListNS = mcp.NewTool( 
	"list-ns",
	mcp.WithDescription("List the namespace in the kubernetes cluster with status"),
)

var GetNS = mcp.NewTool( 
	"get-ns",
	mcp.WithDescription("Get the particular namespace in the kubernetes cluster with status"),
	mcp.WithString(
		"name",
		mcp.Required(),
        mcp.Description("The name of the namespace to get details for"),
	),
)

var DeleteNS = mcp.NewTool( 
	"delete-ns",
	mcp.WithDescription("Delete the particular namespace in the kubernetes cluster"),
	mcp.WithString(
		"name",
		mcp.Required(),
        mcp.Description("The name of the namespace to be deleted"),
	),
)

var UpdateNS = mcp.NewTool(
	"update-ns",
	mcp.WithDescription("Update the namespace like label and annotation changes"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the namespace to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be updated"),
	),
	mcp.WithString(
		"annotation",
		mcp.Description("annotation to be updated"),
	),
)

var CreateNS = mcp.NewTool(
	"create-ns",
	mcp.WithDescription("Create the namespace"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the namespace to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be add in hte namespace"),
	),
)

var ListDeploymentInNamespace = mcp.NewTool(
	"list-deployment-in-namespace",
	mcp.WithDescription("List the deployment in particular namespace with available instance and label"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the deployment should be listed"),
	),
	mcp.WithString(
		"label", 
		mcp.Description("The deployment should be listed only if this particular label is exist"),
	),
)

var ListDeployment = mcp.NewTool(
	"list-deployment",
	mcp.WithDescription("List the deployment in the all namespaces with available instance with label"),
	mcp.WithString(
		"label", 
		mcp.Description("The deployment should be listed only if this particular label is exist"),
	),
)

var GetDeployment = mcp.NewTool(
	"get-deployment",
	mcp.WithDescription("Get the deployment in particular namespace with available instance and label"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace to get the deployment"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the deployment to get"),
	),
)

var DeleteDeployment = mcp.NewTool(
	"delete-deployment",
	mcp.WithDescription("Delete the deployment in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the deployment to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the deployment to be deleted"),
	),
)

var UpdateDeployment = mcp.NewTool(
	"update-deployment",
	mcp.WithDescription("Update the deployment in particular namespace like label, annotation, replica and image changes"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the deployment to be updated"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the deployment to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be updated"),
	),
	mcp.WithString(
		"annotation",
		mcp.Description("annotation to be updated"),
	),
	mcp.WithNumber(
		"replica",
		mcp.Description("Replica to be updated"),
	),
	mcp.WithString(
		"containerName",
		mcp.Description("Container Name to update the image"),
	),
	mcp.WithString(
		"image",
		mcp.Description("Image to be updated"),
	),

)

var CreateDeployment = mcp.NewTool(
	"create-deployment",
	mcp.WithDescription("Create the deployment in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the deployment to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the deployment to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be added in that deployment"),
	),
	mcp.WithNumber(
		"replica",
		mcp.Description("Number of replica"),
	),		
	mcp.WithString(
		"containerNames",
		mcp.Required(),
		mcp.Description("Container Names for the deployment"),
	),
	mcp.WithString(
		"containerImages",
		mcp.Required(),
		mcp.Description("Container Image for the deployment"),
	),
	mcp.WithString(
		"containerPorts",
		mcp.Description("Container port details for the deployment"),
	),

)

var ListServiceInNamespace = mcp.NewTool(
	"list-service-in-namespace",
	mcp.WithDescription("List the service in particular namespace with type"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the service should be listed"),
	),
)

var ListService = mcp.NewTool(
	"list-service",
	mcp.WithDescription("List the service in the all namespace with type"),
)

var GetService = mcp.NewTool(
	"get-service",
	mcp.WithDescription("Get the particular service with type and IP"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the service should be listed"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the service to get"),
	),
)

var DeleteService = mcp.NewTool(
	"delete-service",
	mcp.WithDescription("Delete the particular service in the namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the service to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the service to be deleted"),
	),
)

var UpdateService = mcp.NewTool(
	"update-service",
	mcp.WithDescription("Update the service in particular namespace like selector label and type changes"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the service to be updated"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the service to be updated"),
	),
	mcp.WithString(
		"selectorLabel",
		mcp.Description("Selector label to be updated"),
	),
	mcp.WithString(
		"svctype",
		mcp.Description("Service type to be updated"),
	),
)

var CreateService = mcp.NewTool(
	"create-service",
	mcp.WithDescription("Create the service in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the service to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the service to be created"),
	),
	mcp.WithString(
		"selectorLabel",
		mcp.Required(),
		mcp.Description("Selector Label for the service"),
	),
	mcp.WithString(
		"svcPort",
		mcp.Required(),
		mcp.Description("Service port name and port details for service"),
	),		
	mcp.WithString(
		"targetPort",
		mcp.Required(),
		mcp.Description("Target port details"),
	),
	mcp.WithString(
		"svcType",
		mcp.Description("Service type need to create, if not provided it will take default service type"),
	),
)

var  ListStatefulsetInNamespace = mcp.NewTool(
	"list-statefulset-in-namespace",
	mcp.WithDescription("List the statefulset in particular namespace with available instance and label"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the statefulset should be listed"),
	),
	mcp.WithString(
		"label", 
		mcp.Description("Get the statefulset only if this particular label is exist"),
	),
)

var ListStatefulset = mcp.NewTool(
	"list-statefulset",
	mcp.WithDescription("List the statefulset in the all namespace with available instance and label"),
	mcp.WithString(
		"label", 
		mcp.Description("Get the statefulset only if this particular label is exist"),
	),
)

var  GetStatefulset = mcp.NewTool(
	"get-statefulset",
	mcp.WithDescription("Get the particular statefulset in particular namespace with available instance and labels"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the statefulset should be listed"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the statefulset to get"),
	),
)

var  DeleteStatefulset = mcp.NewTool(
	"delete-statefulset",
	mcp.WithDescription("Delete the particular statefulset in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the statefulset to be deletes"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the statefulset to be deleted"),
	),
)

var UpdateStatefulset = mcp.NewTool(
	"update-statefulset",
	mcp.WithDescription("Update the statefulset in particular namespace like label, annotation, replica and image changes"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the statefulset to be updated"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the statefulset to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be updated"),
	),
	mcp.WithString(
		"annotation",
		mcp.Description("annotation to be updated"),
	),
	mcp.WithNumber(
		"replica",
		mcp.Description("Replica to be updated"),
	),
	mcp.WithString(
		"containerName",
		mcp.Description("Container Name to update the image"),
	),
	mcp.WithString(
		"image",
		mcp.Description("Image to be updated"),
	),
)

var CreateStatefulset = mcp.NewTool(
	"create-statefulset",
	mcp.WithDescription("Create the statefulset in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the statefulset to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the statefulset to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be added in that statefulset"),
	),
	mcp.WithString(
		"containerNames",
		mcp.Description("Container Names for the statefulset"),
	),
	mcp.WithString(
		"containerImages",
		mcp.Required(),
		mcp.Description("Container Image for the statefulset"),
	),
	mcp.WithNumber(
		"containerPorts",
		mcp.Description("Container port for the statefulset"),
	),
	mcp.WithString(
		"storageValue",
		mcp.Required(),
		mcp.Description("Pvc size for the statefulset"),
	),
	mcp.WithString(
		"mountPath",
		mcp.Required(),
		mcp.Description("mount path for the statefulset container to mount the pvc"),
	),
	mcp.WithString(
		"pvcName",
		mcp.Description("Name of the pvc for statefulset"),
	),
	mcp.WithString(
		"svcType",
		mcp.Description("Servcie type for statefulset service"),
	),
	mcp.WithNumber(
		"svcPort",
		mcp.Description("Service Port for the statefulset service"),
	),
	mcp.WithNumber(
		"replica",
		mcp.Description("Number of replica for statefulset"),
	),
)

var ListDaemonsetInNamespace = mcp.NewTool(
	"list-daemonset-in-namespace",
	mcp.WithDescription("List the daemonset in particular namespace with available instance and label"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the daemonset should be listed"),
	),
	mcp.WithString(
		"label", 
		mcp.Description("The daemonset should be listed only if this particular label is exist"),
	),
)

var ListDaemonset = mcp.NewTool(
	"list-daemonset",
	mcp.WithDescription("List the daemonset in the all namespace with available instance and label"),
	mcp.WithString(
		"label", 
		mcp.Description("Get the daemonset only if this particular label is exist"),
	),
)

var GetDaemonset = mcp.NewTool(
	"get-daemonset",
	mcp.WithDescription("Get the daemonset in particular namespace with available instance and label"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace to get the daemonset"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the daemonset to get"),
	),
)

var DeleteDaemonset = mcp.NewTool(
	"delete-daemonset",
	mcp.WithDescription("Delete the daemonset in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the daemonset to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the daemonset to be deleted"),
	),
)

var UpdateDaemonset = mcp.NewTool(
	"update-daemonset",
	mcp.WithDescription("Update the daemonset in particular namespace like label, annotation and image changes"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the daemonset to be updated"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the daemonset to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be updated"),
	),
	mcp.WithString(
		"annotation",
		mcp.Description("annotation to be updated"),
	),
	mcp.WithString(
		"containerName",
		mcp.Description("Container Name to update the image"),
	),
	mcp.WithString(
		"image",
		mcp.Description("Image to be updated"),
	),

)

var CreateDaemonset = mcp.NewTool(
	"create-daemonset",
	mcp.WithDescription("Create the daemonset in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in whcih the daemonset to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the daemonset to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label to be added in that daemonset"),
	),
	mcp.WithString(
		"containerNames",
		mcp.Required(),
		mcp.Description("Container Names for the daemonset"),
	),
	mcp.WithString(
		"containerImages",
		mcp.Required(),
		mcp.Description("Container Image for the daemonset"),
	),
	mcp.WithString(
		"containerPorts",
		mcp.Description("Container port details for the daemonset"),
	),

)

var ListConfigmapInNamespace = mcp.NewTool(
	"list-configmap-in-namespace",
	mcp.WithDescription("List the configmap in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the configmap should be listed"),
	),
)

var ListConfigmap = mcp.NewTool(
	"list-configmap",
	mcp.WithDescription("List the configmap in the all namespace"),
)

var GetConfigmap = mcp.NewTool(
	"get-configmap",
	mcp.WithDescription("Get the configmap in particular namespace with data"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the configmap to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the configmap to get"),
	),
)

var DeleteConfigmap = mcp.NewTool(
	"delete-configmap",
	mcp.WithDescription("Delete the configmap in particular namespace "),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the configmap to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the configmap to be deleted"),
	),
)

var CreateConfigmap = mcp.NewTool(
	"create-configmap",
	mcp.WithDescription("Create the configmap in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the configmap to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the configmap to be created"),
	),
	mcp.WithString(
		"data",
		mcp.Required(),
		mcp.Description("Data of the configmap to be created for"),
	),
)


var ListSecretInNamespace = mcp.NewTool(
	"list-secret-in-namespace",
	mcp.WithDescription("List the secret in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the secret should be listed"),
	),
)

var ListSecret = mcp.NewTool(
	"list-secret",
	mcp.WithDescription("List the secret in the all namespace"),
)

var GetSecret = mcp.NewTool(
	"get-secret",
	mcp.WithDescription("Get the secret in particular namespace with data"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the secret to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the secret to get"),
	),
)

var DeleteSecret = mcp.NewTool(
	"delete-secret",
	mcp.WithDescription("Delete the secret in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the secret to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the secret to be deleted"),
	),
)

var CreateSecret = mcp.NewTool(
	"create-secret",
	mcp.WithDescription("Create the secret in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("The namespace in which the secret to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the secret to be created"),
	),
	mcp.WithString(
		"data",
		mcp.Required(),
		mcp.Description("Data of the secret to be created for"),
	),
)
	
var ListNode = mcp.NewTool(
	"list-node",
	mcp.WithDescription("List the node in the kubernetes cluster with status"),
)

var GetNode = mcp.NewTool(
	"get-node",
	mcp.WithDescription("Get the particular node in the kubernetes cluster with status, kubernetes version, architecture and os"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the node to get"),
	),
)

var DeleteNode = mcp.NewTool(
	"delete-node",
	mcp.WithDescription("Delete the particular node in the kubernetes cluster"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the node to be deleted"),
	),
)

var UpdateNode = mcp.NewTool(
	"update-node",
	mcp.WithDescription("Update the node like label changes"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the node to be updated"),
	),
	mcp.WithString(
		"label",
		mcp.Required(),
		mcp.Description("Label to be updated"),
	),
)

var ListSA = mcp.NewTool(
	"list-serviceAccount",
	mcp.WithDescription("List the serviceAccount in all the namespace"),
	mcp.WithString(
		"label",
		mcp.Description("Label of the serviceAccount, if we need to list the service account with particualr label exist"),
	),
)

var ListSAInNS = mcp.NewTool(
	"list-serviceAccount-in-namepsace",
	mcp.WithDescription("List the serviceAccount in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the serviceAccount to be listed"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label of the serviceAccount, if we need to list the service account with particualr label"),
	),
)

var GetSA = mcp.NewTool(
	"get-serviceAccount",
	mcp.WithDescription("Get the serviceAccount in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the serviceAccount to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the serviceAccount to get"),
	),
)

var DeleteSA = mcp.NewTool(
	"delete-serviceAccount",
	mcp.WithDescription("Delete the serviceAccount in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the serviceAccount to delete"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the serviceAccount to delete"),
	),
)

var CreateSA = mcp.NewTool(
	"create-serviceAccount",
	mcp.WithDescription("Create the serviceAccount in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the serviceAccount to be created"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the serviceAccount to be created"),
	),
	mcp.WithString(
		"label",
		mcp.Description("Label of the serviceAccount, if we need to create the service account with particualr label"),
	),
)

var ListPVCInNS = mcp.NewTool(
	"list-pvc-in-namespace",
	mcp.WithDescription("List the pvc in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the pvc to be listed"),
	),
)

var ListPVC = mcp.NewTool(
	"list-pvc",
	mcp.WithDescription("List the pvc in all namespace"),
)

var GetPVC = mcp.NewTool(
	"get-pvc",
	mcp.WithDescription("Get the pvc in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the pvc to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pvc to get"),
	),
)

var DeletePVC = mcp.NewTool(
	"delete-pvc",
	mcp.WithDescription("Delete the pvc in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace in which the pvc to be deleted"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pvc to delete"),
	),
)

var UpdatePVC = mcp.NewTool(
	"update-pvc",
	mcp.WithDescription("update the pvc in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the pvc to update"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pvc to update"),
	),
	mcp.WithString(
		"size",
		mcp.Required(),
		mcp.Description("size of the pvc to update"),
	),
)

var CreatePVC = mcp.NewTool(
	"create-pvc",
	mcp.WithDescription("Create the pvc in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the pvc to create"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pvc to create"),
	),
	mcp.WithString(
		"size",
		mcp.Required(),
		mcp.Description("Size of the pvc to create"),
	),
	mcp.WithString(
		"storageClass",
		mcp.Required(),
		mcp.Description("Name of the storageClass for pvc to create"),
	),
	mcp.WithString(
		"accessMode",
		mcp.Description("AccessModes of the pvc to create"),
	),
)

var ListPV = mcp.NewTool(
	"list-pv",
	mcp.WithDescription("List the entire pv"),
)

var GetPV = mcp.NewTool(
	"get-pv",
	mcp.WithDescription("Get the pv in particular name"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pv to get"),
	),
)

var DeletePV = mcp.NewTool(
	"delete-pv",
	mcp.WithDescription("Delete the particular pv"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the pv to delete"),
	),
)


var ListRoleInNS = mcp.NewTool(
	"list-role-in-namespace",
	mcp.WithDescription("List the role in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the role to list"),
	),
)

var ListRole = mcp.NewTool(
	"list-role",
	mcp.WithDescription("List the role in all namespace"),
)

var GetRole = mcp.NewTool(
	"get-role",
	mcp.WithDescription("Get the role in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the role to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the role to get"),
	),
)

var ListRBInNS = mcp.NewTool(
	"list-rolebinding-in-namespace",
	mcp.WithDescription("List the rolebinding in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the rolebinding to list"),
	),
)

var ListRB = mcp.NewTool(
	"list-rolebinding",
	mcp.WithDescription("List the rolebinding in all namespace"),
)

var GetRB = mcp.NewTool(
	"get-rolebinding",
	mcp.WithDescription("Get the rolebinding in particular namespace"),
	mcp.WithString(
		"namespace",
		mcp.Required(),
		mcp.Description("Namespace of the rolebinding to get"),
	),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the rolebinding to get"),
	),
)


var ListCR = mcp.NewTool(
	"list-clusterrole",
	mcp.WithDescription("List all the clusterrole in the cluster"),
)

var GetCR = mcp.NewTool(
	"get-clusterrole",
	mcp.WithDescription("Get the particular clusterrole"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the clusterrole to get"),
	),
)

var ListCRB = mcp.NewTool(
	"list-clusterrolebinding",
	mcp.WithDescription("List all the clusterrolebinding in the cluster"),
)

var GetCRB = mcp.NewTool(
	"get-clusterrolebinding",
	mcp.WithDescription("Get the particular clusterrolebinding"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the clusterrolebinding to get"),
	),
)

var ListSC = mcp.NewTool(
	"list-storageClass",
	mcp.WithDescription("List the storageClass in the entier cluster"),
)

var GetSC = mcp.NewTool(
	"get-storageClass",
	mcp.WithDescription("Get the particular storafeClass"),
	mcp.WithString(
		"name",
		mcp.Required(),
		mcp.Description("Name of the storageClass to get"),
	),
)