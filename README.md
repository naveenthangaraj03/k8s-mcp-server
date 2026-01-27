# K8S MCP SERVER

k8s-mcp-server is a Golang based Model Context Protocol (MCP) server that expose Kubernetes resources as structured  MCP tools, enable AI agents to safely interact with Kubernetes cluster.

## Features

### The following kubernetes resources are supported with their respective operations:

- Pod: Create, Get, List, Update, Delete and Log.
- Deployment: Create, Get, List, Update and Delete.
- Daemonset: Create, Get, List, Update and Delete.
- Statefulset: Create, Get, List, Update and Delete.
- Namespace: Create, Get, List, Update and Delete.
- Service: Create, Get, List, Update and Delete.
- Configmap: Create, Get, List and Delete.
- Secret: Create, Get, List and Delete.
- Node: Create, Get, List and Delete.
- ServiceAccount: Create, Get, List and Delete.
- PVC: Create, Get, List, Update and Delete.
- PV: List, Get and Delete.
- Role: Create, Get, List and Delete.
- RoleBinding: Create, Get, List and Delete.
- ClusterRole: Create, Get, List and Delete.
- ClusterRoleBinding: Create, Get, List and Delete.
- Storageclass: Create, Get, List and Delete.
- CRD: Create, Get, List and Delete.

Create supports json data as well.

All interactions are performed via Kubernetes API using the provided kubeconfig.

### Support Custom Resource:

In addition to predefined kubernetes resource tools listed above, this MCP Server also provides a generic custom tool that allows user to interact with any kubernetes custom resource that is not explicitly supported.

This is useful when:

- You want to work with custom resource(CRDs).
- You want to access new kubernetes resources without updating MCP Server.

The MCP Server forwards the provided parameters to the gRPC backend server, which dynamically resolves the resource and perform the requested action.

Custom Tool Supported Operations: Create, Get, List and Delete.


NOTE: Parameter details for each resource are available in the respective `README.md` file under the kubernetes directory.
 
## Prerequisites

- Go
- Access to kubernetes cluster
- Kubeconfig file
- An application with MCP supported

## Installation

```
go install github.com/naveenthangaraj03/k8s-mcp-server@latest
```

## Running MCP Server

Claude Desktop:
Add the following configuration to yours claude config file.

```
{
    "mcpServers": {
        "Kubernetes": {
            "command": "k8s-mcp-server",
            "args": ["--kubeconfigPath=<Path to kubeconfig file>"]
        }
    }
}
```

Enable the custom tool by using the `--customURL` flag.

```
{
    "mcpServers": {
        "Kubernetes": {
            "command": "k8s-mcp-server",
            "args": ["--kubeconfigPath=<Path to kubeconfig file>","--customURL=<grpc server url>"]
        }
    }
}
```

## Security Concern

- Access is fully controlled by the RBAC permisiion defined in the kubeconfig.
- Only operation allowed by the kubeconfig is executed.
