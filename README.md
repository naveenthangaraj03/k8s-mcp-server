# K8S MCP SERVER

k8s-mcp-server is a Golang based Model Context Protocol (MCP) server that expose Kubernetes resources as structured  MCP tools, enable AI agents to safely interact with Kubernetes cluster.

### Features

The following kubernetes resources are supported with their respective operations:

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
- Role: Get and List.
- RoleBinding: Get and List.
- ClusterRole: Get and List.
- ClusterRoleBinding: Get and List.
- Storageclass: Get and List.

All interactions are performed via Kubernetes API using the provided kubeconfig.

### Prerequisites

- Go
- Access to kubernetes cluster
- Kubeconfig file
- An application with MCP supported

### Installation

```
go install github.com/naveenthangaraj03/k8s-mcp-server@latest
```

### Running MCP Server

Claude Desktop:
Add the following configuration to yours claude config file.

```
{
    "mcpServers": {
        "Kubernetes": {
            "command": "k8s-mcp-server",
            "args": [ "<Path to kubeconfig file>"]
        }
    }
}
```

### Security Concern

- Access is fully controlled by the RBAC permisiion defined in the kubeconfig.
- Only operation allowed by the kubeconfig is executed.