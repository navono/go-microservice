package service

import (
	consul "github.com/hashicorp/consul/api"
)

func (s *Service) consulRegister() {
	cfg := consul.DefaultConfig()
	cfg.Address = "192.168.200.110:8500"
	c, err := consul.NewClient(cfg)
	if err != nil {
		panic("Failed to connect to Consul agent")
	}
	s.ConsulAgent = c.Agent()

	serviceDef := &consul.AgentServiceRegistration{
		Address: "192.168.200.100",
		Name:    s.Name,
		Port:    s.Port,
		// Check: &consul.AgentServiceCheck{
		// 	TTL: s.TTL.String(),
		// },
		Check: &consul.AgentServiceCheck{
			// HTTP:     "http://192.168.200.100:8989/health",
			HTTP:     "http://localhost:8989/health",
			Interval: "5s",
		},
	}

	if err := s.ConsulAgent.ServiceRegister(serviceDef); err != nil {
		panic("Failed to register Service")
	}
}
