package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

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
		fmt.Println(field("ContainerID:"), v.ID)
		fmt.Println(field("Created:"), v.Created)
		fmt.Println(field("Status:"), v.State.Status)

		fmt.Println("---------------------------------------")
	}
}

func field(s string) string {
	return fmt.Sprintf("%-15s", s)
}
