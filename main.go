package main

import (
	"fmt"
	"flag"
	"github.com/mark3labs/mcp-go/server"
	"github.com/naveenthangaraj03/k8s-mcp-server/tools"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pod"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/namespace"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/deployment"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/daemonset"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/statefulset"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/service"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/node"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/configmap"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/secret"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/serviceaccount"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/role"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/rolebinding"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pvc"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/pv"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/clusterrole"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/clusterrolebinding"
	"github.com/naveenthangaraj03/k8s-mcp-server/kubernetes/storageclass"
)


func main() {
	s := server.NewMCPServer(
		"Kubernetes MCP",
        "1.0.0",
	)

	flag.Parse()

	s.AddTool(tools.ListPodInNamespace, pod.ListPodInNS)
	s.AddTool(tools.ListPod, pod.ListPod)
	s.AddTool(tools.GetPod, pod.GetPod)
	s.AddTool(tools.DeletePod, pod.DeletePod)
	s.AddTool(tools.UpdatePod, pod.UpdatePod)
	s.AddTool(tools.CreatePod, pod.CreatePod)
	s.AddTool(tools.PodLog, pod.PodLog)


	s.AddTool(tools.ListNS, namespace.ListNS)
	s.AddTool(tools.GetNS, namespace.GetNS)
	s.AddTool(tools.DeleteNS, namespace.DeleteNS)
	s.AddTool(tools.UpdateNS, namespace.UpdateNS)
	s.AddTool(tools.CreateNS, namespace.CreateNS)


	s.AddTool(tools.ListNode, node.ListNode)
	s.AddTool(tools.GetNode, node.GetNode)
	s.AddTool(tools.DeleteNode, node.DeleteNode)
	s.AddTool(tools.UpdateNode, node.UpdateNode)


	s.AddTool(tools.ListDeploymentInNamespace, deployment.ListDeploymentInNS)
	s.AddTool(tools.ListDeployment, deployment.ListDeployment)
	s.AddTool(tools.GetDeployment, deployment.GetDeployment)
	s.AddTool(tools.DeleteDeployment, deployment.DeleteDeployment)
	s.AddTool(tools.CreateDeployment, deployment.CreateDeployment)
	s.AddTool(tools.UpdateDeployment, deployment.UpdateDeployment)


	s.AddTool(tools.ListDaemonsetInNamespace, daemonset.ListDaemonsetInNS)
	s.AddTool(tools.ListDaemonset, daemonset.ListDaemonset)
	s.AddTool(tools.GetDaemonset, daemonset.GetDaemonset)
	s.AddTool(tools.DeleteDaemonset, daemonset.DeleteDaemonset)
	s.AddTool(tools.UpdateDaemonset, daemonset.UpdateDaemonset)
	s.AddTool(tools.CreateDaemonset, daemonset.CreateDaemonset)


	s.AddTool(tools.ListStatefulsetInNamespace, statefulset.ListStatefulsetInNS)
	s.AddTool(tools.ListStatefulset, statefulset.ListStatefulset)
	s.AddTool(tools.GetStatefulset, statefulset.GetStatefulset)
	s.AddTool(tools.DeleteStatefulset, statefulset.DeleteStatefulset)
	s.AddTool(tools.UpdateStatefulset, statefulset.UpdateStatefulset)
	s.AddTool(tools.CreateStatefulset, statefulset.CreateStatefulset)


	s.AddTool(tools.ListServiceInNamespace, service.ListServiceInNS)
	s.AddTool(tools.ListService, service.ListService)
	s.AddTool(tools.GetService, service.GetService)
	s.AddTool(tools.DeleteService, service.GetService)
	s.AddTool(tools.UpdateService, service.UpdateService)
	s.AddTool(tools.CreateService, service.CreateService)


	s.AddTool(tools.ListConfigmapInNamespace, configmap.ListConfigmapInNS)
	s.AddTool(tools.ListConfigmap, configmap.ListConfigmap)
	s.AddTool(tools.GetConfigmap, configmap.GetConfigmap)
	s.AddTool(tools.DeleteConfigmap, configmap.DeleteConfigmap)
	s.AddTool(tools.CreateConfigmap, configmap.CreateConfigmap)


	s.AddTool(tools.ListSecretInNamespace, secret.ListSecretInNS)
	s.AddTool(tools.ListSecret, secret.ListSecret)
	s.AddTool(tools.GetSecret, secret.GetSecret)
	s.AddTool(tools.DeleteSecret, secret.DeleteSecret)
	s.AddTool(tools.CreateSecret, secret.CreateSecret)
	
	
	s.AddTool(tools.ListSA, serviceaccount.ListSA)
	s.AddTool(tools.ListSAInNS, serviceaccount.ListSAInNS)
	s.AddTool(tools.GetSA, serviceaccount.GetSA)
	s.AddTool(tools.DeleteSA, serviceaccount.DeleteSA)
	s.AddTool(tools.CreateSA, serviceaccount.CreateSA)

	s.AddTool(tools.ListRole, role.ListRole)
	s.AddTool(tools.ListRoleInNS, role.ListRoleInNS)
	s.AddTool(tools.GetRole, role.GetRole)
	
	s.AddTool(tools.ListRB, rolebinding.ListRB)
	s.AddTool(tools.ListRBInNS, rolebinding.ListRBInNS)
	s.AddTool(tools.GetRB, rolebinding.GetRB)

	s.AddTool(tools.ListPVC, pvc.ListPVC)
	s.AddTool(tools.ListPVCInNS, pvc.ListPVCInNS)
	s.AddTool(tools.GetPVC, pvc.GetPVC)
	s.AddTool(tools.DeletePVC, pvc.DeletePVC)
	s.AddTool(tools.UpdatePVC, pvc.UpdatePVC)

	s.AddTool(tools.ListPV, pv.ListPV)
	s.AddTool(tools.GetPV, pv.GetPV)
	s.AddTool(tools.DeletePV, pv.DeletePV)

	s.AddTool(tools.ListCR, clusterrole.ListCR)
	s.AddTool(tools.GetCR, clusterrole.GetCR)

	s.AddTool(tools.ListCRB, clusterrolebinding.ListCRB)
	s.AddTool(tools.GetCRB, clusterrolebinding.GetCRB)

	s.AddTool(tools.ListSC, storageclass.ListSC)
	s.AddTool(tools.GetSC, storageclass.GetSC)

    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Error starting server: %v\n", err)
    }
}
