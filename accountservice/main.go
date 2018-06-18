package main

import (
	"fmt"

	"github.com/navono/go-microservice/accountservice/service"
)

var appName = "accountservice"

func main() {
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("8989")
}
