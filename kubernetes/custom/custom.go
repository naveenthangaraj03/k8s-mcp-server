package custom

import (
	"context"
	"flag"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	rpc "github.com/naveenthangaraj03/k8s-mcp-server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var customURL string

func init() {
	flag.StringVar(&customURL, "customURL", "", "Custom URL for custom tool")
}

func Custom(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	kind, err := request.RequireString("kind")
	if err != nil {
		output := fmt.Sprintf("Provide kind for custom resource")
		return mcp.NewToolResultText(string(output)), nil
	}
	method, err := request.RequireString("method")
	if err != nil {
		output := fmt.Sprintf("Provide method to do action")
		return mcp.NewToolResultText(string(output)), nil
	}
	name := request.GetString("name", "")
	namespace := request.GetString("namespace", "")
	jsondata := request.GetString("jsondata", "")

	if customURL == "" {
		output := fmt.Sprintf("Provide custom URL to connect to grpc server")
		return mcp.NewToolResultText(string(output)), nil
	}

	conn, err := grpc.Dial(customURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		output := fmt.Sprintf("Failed to connect grpc server:", err)
		return mcp.NewToolResultText(string(output)), nil
	}

	client := rpc.NewCustomToolServiceClient(conn)

	req := &rpc.CustomRequest{
		Kind:      kind,
		Method:    method,
		Name:      name,
		Namespace: namespace,
		JsonData:  jsondata,
	}

	res, err := client.CustomTool(context.Background(), req)
	if err != nil {
		output := fmt.Sprintf("Failed to get response from server:", err)
		return mcp.NewToolResultText(string(output)), nil
	}

	output := res.Result

	return mcp.NewToolResultText(string(output)), nil
}
