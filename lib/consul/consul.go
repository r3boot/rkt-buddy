package consul

import (
	"fmt"
	"net/http"
	"os"

	"time"

	"strings"

	"github.com/r3boot/rkt-buddy/lib/config"
	"github.com/r3boot/rkt-buddy/lib/logger"
)

var (
	log *logger.Logger
)

func NewConsul(l *logger.Logger, cfg *config.Config) (*Consul, error) {
	log = l

	endpoint := "http://localhost:8500"
	if value := os.Getenv("CONSUL_ENDPOINT"); value != "" {
		endpoint = ""
	}

	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    5 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{
		Transport: transport,
	}

	agent := &Consul{
		config:       cfg,
		endpoint:     endpoint,
		client:       client,
		HealthEvents: make(chan bool),
	}

	clusterMembers, err := agent.GetMembers()
	if err != nil {
		log.Warningf("NewConsul: Failed to get cluster clusterMembers, disabling consul support")
		return agent, nil
	}

	allClusterMembers := []string{}
	for _, entry := range *clusterMembers {
		allClusterMembers = append(allClusterMembers, entry.Name)
	}

	log.Debugf("NewConsul: Found %d cluster nodes: %s", len(allClusterMembers), strings.Join(allClusterMembers, ", "))

	svcName := os.Getenv("AC_APP_NAME")
	if svcName == "" {
		return nil, fmt.Errorf("Buddy.FetchServiceMeta: AC_APP_NAME not set")
	}
	svcInstance := fmt.Sprintf("%s-%s", svcName, cfg.Service.Node)

	svcDescr := os.Getenv("BUDDY_SVC_DESC")

	agent.registerMeta = RegisterData{
		Datacenter: cfg.Service.Datacenter,
		Id:         GenUuid(),
		Node:       cfg.Service.Node,
		Address:    cfg.Service.Address,
		Service: ServiceData{
			Id:      svcInstance,
			Service: svcName,
			Address: cfg.Service.Address,
			Port:    cfg.Service.Port,
		},
		Check: CheckData{
			Node:      cfg.Service.Node,
			CheckId:   fmt.Sprintf("service:%s", svcName),
			Name:      fmt.Sprintf("Health check for %s", svcName),
			Notes:     svcDescr,
			ServiceId: cfg.Service.Instance,
			Definition: CheckDefinitionData{
				Http:     fmt.Sprintf("http://%s:6548/v1/health", cfg.Service.Address),
				Interval: "5s",
				Timeout:  "1s",
				DeregisterCriticalServiceAfter: "30s",
			},
		},
	}

	agent.deregisterMeta = DeregisterData{
		Datacenter: cfg.Service.Datacenter,
		Node:       cfg.Service.Node,
	}

	return agent, nil
}
