package service

import (
	"fmt"
	"log"
	"net/http"

	consul "github.com/hashicorp/consul/api"
)

// Service for service type
type Service struct {
	Name string
	Port int
	// TTL         time.Duration
	// RedisClient redis.UniversalClient
	ConsulAgent *consul.Agent
	// Metrics     Metrics
}

// Metrics for prometheus
// type Metrics struct {
// 	RedisRequests  *prometheus.CounterVec
// 	RedisDurations prometheus.Summary
// }

// New create a new Service instance
func New(port int) (*Service, error) {
	s := new(Service)
	s.Name = "account"
	s.Port = port
	s.consulRegister()

	return s, nil
}

// Start start a HTTP server
func (s *Service) Start() {
	r := NewRouter()
	http.Handle("/", r)

	port := fmt.Sprintf("%d", s.Port)
	log.Println("Starting HTTP service at " + port)

	err := http.ListenAndServe("0.0.0.0:"+port, nil)
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}
