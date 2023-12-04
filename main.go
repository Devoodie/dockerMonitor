package main

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"net"
	"net/http"
	"net/rpc"
)

func dockerps() []types.Container {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}
func main() {
	test := 1
	rpc.Register(test)
	rpc.HandleHTTP()
	line, err := net.Listen("tcp", ":45000")
	if err != nil {
		panic(err)
	}
	go func() {
		err := http.Serve(line, nil)
		if err != nil {
			panic(err)
		}
	}()
}
