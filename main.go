package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"github.com/mark3labs/mcp-go/server"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/clusterrole"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/clusterrolebinding"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/configmap"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/crd"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/custom"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/daemonset"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/deployment"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/namespace"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/node"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pod"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pv"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pvc"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/role"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/rolebinding"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/secret"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/service"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/serviceaccount"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/statefulset"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/storageclass"
	"github.com/naveenthangaraj03/k8s-mcp-server/tools"
)

var mode string

func init() {
	flag.StringVar(&mode, "mode", "stdio", "MCP server mode")
}

func main() {
	s := server.NewMCPServer(
		"Kubernetes MCP",
		"1.0.0",
	)

	flag.Parse()

	// Pod tools
	s.AddTool(tools.ListPodInNamespace, pod.ListPodInNS)
	s.AddTool(tools.ListPod, pod.ListPod)
	s.AddTool(tools.GetPod, pod.GetPod)
	s.AddTool(tools.DeletePod, pod.DeletePod)
	s.AddTool(tools.UpdatePod, pod.UpdatePod)
	s.AddTool(tools.CreatePod, pod.CreatePod)
	s.AddTool(tools.PodLog, pod.PodLog)
	s.AddTool(tools.CreatePodWithJson, pod.CreatePodWithJson)

	// Namespace tools
	s.AddTool(tools.ListNS, namespace.ListNS)
	s.AddTool(tools.GetNS, namespace.GetNS)
	s.AddTool(tools.DeleteNS, namespace.DeleteNS)
	s.AddTool(tools.UpdateNS, namespace.UpdateNS)
	s.AddTool(tools.CreateNS, namespace.CreateNS)
	s.AddTool(tools.CreateNSWithJson, namespace.CreateNSWithJson)

	// Node tools
	s.AddTool(tools.ListNode, node.ListNode)
	s.AddTool(tools.GetNode, node.GetNode)
	s.AddTool(tools.DeleteNode, node.DeleteNode)
	s.AddTool(tools.UpdateNode, node.UpdateNode)

	// Deployment tools
	s.AddTool(tools.ListDeploymentInNamespace, deployment.ListDeploymentInNS)
	s.AddTool(tools.ListDeployment, deployment.ListDeployment)
	s.AddTool(tools.GetDeployment, deployment.GetDeployment)
	s.AddTool(tools.DeleteDeployment, deployment.DeleteDeployment)
	s.AddTool(tools.CreateDeployment, deployment.CreateDeployment)
	s.AddTool(tools.UpdateDeployment, deployment.UpdateDeployment)
	s.AddTool(tools.CreateDeploymentWithJson, deployment.CreateDeploymentWithJson)

	// Daemonset tools
	s.AddTool(tools.ListDaemonsetInNamespace, daemonset.ListDaemonsetInNS)
	s.AddTool(tools.ListDaemonset, daemonset.ListDaemonset)
	s.AddTool(tools.GetDaemonset, daemonset.GetDaemonset)
	s.AddTool(tools.DeleteDaemonset, daemonset.DeleteDaemonset)
	s.AddTool(tools.UpdateDaemonset, daemonset.UpdateDaemonset)
	s.AddTool(tools.CreateDaemonset, daemonset.CreateDaemonset)
	s.AddTool(tools.CreateDaemonsetWithJson, daemonset.CreateDaemonsetWithJson)

	// Statefulset tools
	s.AddTool(tools.ListStatefulsetInNamespace, statefulset.ListStatefulsetInNS)
	s.AddTool(tools.ListStatefulset, statefulset.ListStatefulset)
	s.AddTool(tools.GetStatefulset, statefulset.GetStatefulset)
	s.AddTool(tools.DeleteStatefulset, statefulset.DeleteStatefulset)
	s.AddTool(tools.UpdateStatefulset, statefulset.UpdateStatefulset)
	s.AddTool(tools.CreateStatefulset, statefulset.CreateStatefulset)
	s.AddTool(tools.CreateStatefulsetWithJson, statefulset.CreateStatefulsetWithJson)

	// Service tools
	s.AddTool(tools.ListServiceInNamespace, service.ListServiceInNS)
	s.AddTool(tools.ListService, service.ListService)
	s.AddTool(tools.GetService, service.GetService)
	s.AddTool(tools.DeleteService, service.GetService)
	s.AddTool(tools.UpdateService, service.UpdateService)
	s.AddTool(tools.CreateService, service.CreateService)
	s.AddTool(tools.CreateServiceWithJson, service.CreateServiceWithJson)

	// Configmap tools
	s.AddTool(tools.ListConfigmapInNamespace, configmap.ListConfigmapInNS)
	s.AddTool(tools.ListConfigmap, configmap.ListConfigmap)
	s.AddTool(tools.GetConfigmap, configmap.GetConfigmap)
	s.AddTool(tools.DeleteConfigmap, configmap.DeleteConfigmap)
	s.AddTool(tools.CreateConfigmap, configmap.CreateConfigmap)
	s.AddTool(tools.CreateConfigmapWithJson, configmap.CreateConfigmapWithJson)

	// Secret tools
	s.AddTool(tools.ListSecretInNamespace, secret.ListSecretInNS)
	s.AddTool(tools.ListSecret, secret.ListSecret)
	s.AddTool(tools.GetSecret, secret.GetSecret)
	s.AddTool(tools.DeleteSecret, secret.DeleteSecret)
	s.AddTool(tools.CreateSecret, secret.CreateSecret)
	s.AddTool(tools.CreateSecretWithJson, secret.CreateSecretWithJson)

	// ServiceAccount tools
	s.AddTool(tools.ListSA, serviceaccount.ListSA)
	s.AddTool(tools.ListSAInNS, serviceaccount.ListSAInNS)
	s.AddTool(tools.GetSA, serviceaccount.GetSA)
	s.AddTool(tools.DeleteSA, serviceaccount.DeleteSA)
	s.AddTool(tools.CreateSA, serviceaccount.CreateSA)
	s.AddTool(tools.CreateSAWithJson, serviceaccount.CreateSAWithJson)

	// Role tools
	s.AddTool(tools.ListRole, role.ListRole)
	s.AddTool(tools.ListRoleInNS, role.ListRoleInNS)
	s.AddTool(tools.GetRole, role.GetRole)
	s.AddTool(tools.DeleteRole, role.DeleteRole)
	s.AddTool(tools.CreateRoleWithJson, role.CreateRoleWithJson)

	// RoleBinding tools
	s.AddTool(tools.ListRB, rolebinding.ListRB)
	s.AddTool(tools.ListRBInNS, rolebinding.ListRBInNS)
	s.AddTool(tools.GetRB, rolebinding.GetRB)
	s.AddTool(tools.DeleteRB, rolebinding.DeleteRB)
	s.AddTool(tools.CreateRBWithJson, rolebinding.CreateRBWithJson)

	// PVC tools
	s.AddTool(tools.ListPVC, pvc.ListPVC)
	s.AddTool(tools.ListPVCInNS, pvc.ListPVCInNS)
	s.AddTool(tools.GetPVC, pvc.GetPVC)
	s.AddTool(tools.DeletePVC, pvc.DeletePVC)
	s.AddTool(tools.UpdatePVC, pvc.UpdatePVC)
	s.AddTool(tools.CreatePVC, pvc.CreatePVC)
	s.AddTool(tools.CreatePVCWithJson, pvc.CreatePVCWithJson)

	// PV tools
	s.AddTool(tools.ListPV, pv.ListPV)
	s.AddTool(tools.GetPV, pv.GetPV)
	s.AddTool(tools.DeletePV, pv.DeletePV)

	// ClusterRole tools
	s.AddTool(tools.ListCR, clusterrole.ListCR)
	s.AddTool(tools.GetCR, clusterrole.GetCR)
	s.AddTool(tools.DeleteCR, clusterrole.DeleteCR)
	s.AddTool(tools.CreateCRWithJson, clusterrole.CreateCRWithJson)

	// ClusterRoleBinding tools
	s.AddTool(tools.ListCRB, clusterrolebinding.ListCRB)
	s.AddTool(tools.GetCRB, clusterrolebinding.GetCRB)
	s.AddTool(tools.DeleteCRB, clusterrolebinding.DeleteCRB)
	s.AddTool(tools.CreateCRBWithJson, clusterrolebinding.CreateCRBWithJson)

	// StorageClass tools
	s.AddTool(tools.ListSC, storageclass.ListSC)
	s.AddTool(tools.GetSC, storageclass.GetSC)
	s.AddTool(tools.DeleteSC, storageclass.DeleteSC)
	s.AddTool(tools.CreateSCWithJson, storageclass.CreateSCWithJson)

	// CRD tools
	s.AddTool(tools.ListCRD, crd.ListCRD)
	s.AddTool(tools.GetCRD, crd.GetCRD)
	s.AddTool(tools.DeleteCRD, crd.DeleteCRD)
	s.AddTool(tools.CreateCRDWithJson, crd.CreateCRDWithJson)

	// Custom tool
	s.AddTool(tools.Custom, custom.Custom)

	if mode == "http" {
		handler := server.NewStreamableHTTPServer(s)
		http.HandleFunc("/mcp", handler.ServeHTTP)
		log.Println("Starting Kuberentes MCP Server")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Printf("Error starting http server: %v\n", err)
		}
	} else {
		if err := server.ServeStdio(s); err != nil {
			fmt.Printf("Error starting stdio server: %v\n", err)
		}
	}
}
