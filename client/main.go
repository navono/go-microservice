package main

import (
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

func main() {
	cfg := consul.DefaultConfig()
	cfg.Address = "localhost:8500"
	c, err := consul.NewClient(cfg)
	if err != nil {
		panic("Failed to connect to Consul agent")
	}

	catalog := c.Catalog()

	// datacenters, _ := catalog.Datacenters()
	// fmt.Println(datacenters)

	// nodes, meta, _ := catalog.Nodes(nil)
	// fmt.Println(meta)

	// for node := range nodes {
	// 	fmt.Println(node)
	// }

	service, meta, _ := catalog.Service("account", "", nil)
	fmt.Printf("%v", service[0])
	fmt.Println(meta)
}
