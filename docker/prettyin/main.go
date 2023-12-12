package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/docker/go-connections/nat"
	"github.com/moby/moby/api/types"
)

func main() {
	raw, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	list := []types.ContainerJSON{}

	err = json.Unmarshal(raw, &list)
	if err != nil {
		panic(err)
	}

	for _, v := range list {
		printLine("Basic")
		printField("Image", v.Config.Image)
		printField("ContainerID", v.ID)
		printField("Status", v.State.Status)
		printField("Created", v.Created)
		printField("Started At", v.State.StartedAt)
		printField("Restart Count", v.RestartCount)
		printField("CMD", v.Config.Cmd)
		printField("Entrypoint", v.Config.Entrypoint)
		printList("Env", v.Config.Env)
		printLine("Network")
		for name, network := range v.NetworkSettings.Networks {
			printField(fmt.Sprintf("IPAddress(%s)", name), network.IPAddress)
		}
		for containerPort, bindings := range v.NetworkSettings.Ports {
			printList(fmt.Sprintf("Ports(%s)", containerPort), bindings)
		}
		printLine("Mount")
		printField("Host Name", v.HostnamePath)
		printField("Hosts path", v.HostsPath)
		printField("Log path", v.LogPath)
		printField("Volumes", v.Config.Volumes)

		fmt.Printf("\n\n")
	}
}

func printLine(s string) {
	fmt.Printf("--------------------%s--------------------\n", s)
}

func printField(name string, val any) {
	fmt.Printf("%-18s%v\n", name+":", val)
}

func printList(name string, a any) {
	switch list := a.(type) {
	case []string:
		{
			for key, val := range list {
				var label string
				if key == 0 {
					label = name + ":"
				}
				fmt.Printf("%-18s%v\n", label, val)
			}
		}
	case map[string]struct{}:
		{
			var label = name + ":"
			for key, val := range list {
				fmt.Printf("%-18s%v  %v\n", label, key, val)
				label = ""
			}
		}
	case []nat.PortBinding:
		{
			for key, val := range list {
				var label string
				if key == 0 {
					label = name + ":"
				}
				fmt.Printf("%-18s%v\n", label, val)
			}
		}
	}
}
