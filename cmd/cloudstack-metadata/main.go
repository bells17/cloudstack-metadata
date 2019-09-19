package main

import (
	"encoding/json"
	"fmt"

	"github.com/bells17/cloudstack-metadata/pkg/metadata"
)

func main() {
	m := metadata.NewMetadata("data-server")
	result, err := m.FetchAll()
	if err != nil {
		fmt.Printf("error occured. err: %v\n", err)
		return
	}

	bytes, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("JSON marshal error occred. err: %v\n", err)
		return
	}
	fmt.Println(string(bytes))
}
