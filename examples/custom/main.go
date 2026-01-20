package main

import (
	"context"
	"encoding/json"
	"fmt"
	rpc "github.com/naveenthangaraj03/k8s-mcp-server/proto"
	"google.golang.org/grpc"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
	"net"
)

type Server struct {
	rpc.CustomToolServiceServer
}

type output struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

var kubeconfigPath string

func kafkaList(group, version, resource string, ctx context.Context) ([]output, error) {
	kubeconfigPath = "/root/.kube/config"
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfigPath)
	if err != nil {
		fmt.Println("Clientcmd error:", err)
		return nil, err
	}
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		fmt.Println("Dynamic error:", err)
		return nil, err
	}
	resourceId := schema.GroupVersionResource{
		Group:    group,
		Version:  version,
		Resource: resource,
	}
	ResourceClient := dynClient.Resource(resourceId)
	Resources, err := ResourceClient.List(ctx, metav1.ListOptions{})
	if err != nil {
		fmt.Printf("Error listing %s Resources: %s\n", resource, err.Error())
		return nil, err
	}
	releases := []output{}
	for _, item := range Resources.Items {
		releases = append(releases, output{
			Name:      item.GetName(),
			Namespace: item.GetNamespace(),
		})
	}
	return releases, nil
}

func main() {
	var addr string = ":8080"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Error in listening", err)
		panic(err)
	}

	grpcServer := grpc.NewServer()

	rpc.RegisterCustomToolServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(lis); err != nil {
		fmt.Println("Error in serve", err)
		panic(err)
	}
}

func (s *Server) CustomTool(ctx context.Context, req *rpc.CustomRequest) (*rpc.CustomResponse, error) {
	if req.Method == "list" {
		if req.Kind == "kafka" {
			fmt.Println("Running Custom tool")
			kafkalist, err := kafkaList("kafka.strimzi.io", "v1beta2", "kafkas", ctx)
			if err != nil {
				fmt.Println("Listing error:", err)
				return nil, err
			}

			data, err := json.Marshal(kafkalist)
			if err != nil {
				fmt.Println("marshal error:", err)
				return nil, err
			}

			return &rpc.CustomResponse{Result: string(data)}, nil
		}
	}

	return &rpc.CustomResponse{Result: "Method or Kind is not supported"}, nil
}
