package main

import (
	"context"
	"encoding/json"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"io"
	"log"
	"net/http"
)

func dockerps() []types.Container {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}
	return containers
}

func main() {
	containerdata := func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		containers, _ := json.Marshal(dockerps())
		io.WriteString(w, string(containers))
	}
	http.HandleFunc("/dockerps", containerdata)
	log.Fatal(http.ListenAndServe(":45000", nil))
}
