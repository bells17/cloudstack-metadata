package main

import (
	"fmt"

	"github.com/bells17/cloudstack-metadata/pkg/metadata"
)

func main() {
	m := metadata.NewMetadata("data-server")
	result, err := m.FetchAll()
	if err != nil {
		fmt.Printf("error occured. err: %v\n", err)
	}
	fmt.Printf("result: %v\n", result)
}
