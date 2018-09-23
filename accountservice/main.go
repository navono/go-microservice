package main

import (
	"fmt"

	"github.com/navono/go-microservice/accountservice/dbclient"
	"github.com/navono/go-microservice/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	initializeBoltClient()
	s, _ := service.New(8989)
	s.Start()
	// service.StartWebServer("8989")
}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
