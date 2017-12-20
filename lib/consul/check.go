package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type AgentCheckData struct {
	Id                             string `json:"ID"`
	Name                           string `json:"Name"`
	Notes                          string `json:"Notes"`
	DeregisterCriticalServiceAfter string `json:"DeregisterCriticalServiceAfter"`
	Http                           string `json:"HTTP"`
	Method                         string `json:"Method"`
	Interval                       string `json:"Interval"`
}

func (c *Consul) RegisterCheck() error {

	data, err := json.Marshal(c.registerCheckMeta)
	if err != nil {
		return fmt.Errorf("Consul.RegisterCheck json.Marshal: %v", err)
	}

	log.Debugf("data: %s", string(data))

	endpoint := fmt.Sprintf("%s/v1/agent/check/register", c.endpoint)

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("Consul.RegisterCheck http.NewRequest: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Consul.RegisterCheck c.client.Do: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Consul.RegisterCheck: failed to register check")
	}

	return nil
}
