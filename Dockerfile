FROM ubuntu:22.04

COPY k8s-mcp-server / 

ENTRYPOINT ["./k8s-mcp-server"]


