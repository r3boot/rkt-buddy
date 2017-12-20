package consul

import (
	"fmt"
	"net/http"
	"os"

	"time"

	"strings"

	"strconv"

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
		endpoint = value
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
		runAgent:     true,
		HealthEvents: make(chan bool),
	}

	clusterMembers, err := agent.GetMembers()
	if err != nil {
		agent.runAgent = false
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
		return nil, fmt.Errorf("NewConsul: AC_APP_NAME not set")
	}
	svcInstance := fmt.Sprintf("%s-%s", svcName, cfg.Service.Node)

	svcDescr := os.Getenv("BUDDY_SVC_DESC")

	svcPort := 0
	value := os.Getenv("BUDDY_SVC_PORT")
	if value != "" {
		svcPort, err = strconv.Atoi(value)
		if err != nil {
			return nil, fmt.Errorf("NewConsul strconv.Atoi: %v", err)
		}
	}

	agent.registerCatalogMeta = RegisterData{
		Datacenter: cfg.Service.Datacenter,
		Id:         GenUuid(),
		Node:       cfg.Service.Node,
		Address:    cfg.Service.Address,
		Service: ServiceData{
			Id:      svcInstance,
			Service: svcName,
			Address: cfg.Service.Address,
			Port:    svcPort,
		},
		Check: CheckData{
			Node:      cfg.Service.Node,
			CheckId:   fmt.Sprintf("service:%s", svcName),
			Name:      fmt.Sprintf("Health check for %s", svcName),
			Notes:     svcDescr,
			ServiceId: svcInstance,
			Definition: CheckDefinitionData{
				Http:     fmt.Sprintf("http://%s:6548/v1/health", cfg.Service.Address),
				Method:   http.MethodGet,
				Interval: "5s",
				Timeout:  "1s",
				DeregisterCriticalServiceAfter: "30s",
			},
		},
	}

	agent.registerCheckMeta = AgentCheckData{
		Id:                             svcInstance,
		Name:                           fmt.Sprintf("Health check for %s", svcName),
		Notes:                          svcDescr,
		Http:                           fmt.Sprintf("http://%s:6548/v1/health", cfg.Service.Address),
		Method:                         http.MethodGet,
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "30s",
	}

	agent.deregisterCatalogMeta = DeregisterData{
		Datacenter: cfg.Service.Datacenter,
		Node:       cfg.Service.Node,
	}

	return agent, nil
}
