package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type CheckDefinitionData struct {
	Http                           string `json:"HTTP"`
	Method                         string `json:"Method"`
	Interval                       string `json:"Interval"`
	Timeout                        string `json:"Timeout"`
	DeregisterCriticalServiceAfter string `json:"DeregisterCriticalServiceAfter"`
}

type CheckData struct {
	Node       string              `json:"Node"`
	CheckId    string              `json:"CheckID"`
	Name       string              `json:"Name"`
	Notes      string              `json:"Notes"`
	ServiceId  string              `json:"ServiceID"`
	Definition CheckDefinitionData `json:"Definition"`
}

type ServiceData struct {
	Id      string   `json:"ID"`
	Service string   `json:"Service"`
	Address string   `json:"Address"`
	Port    int      `json:"Port"`
	Tags    []string `json:"Tags"`
}

type RegisterData struct {
	Datacenter     string      `json:"Datacenter"`
	Id             string      `json:"ID"`
	Node           string      `json:"Node"`
	Address        string      `json:"Address"`
	SkipNodeUpdate bool        `json:"SkipNodeUpdate"`
	Service        ServiceData `json:"Service"`
	Check          CheckData   `json:"Check"`
}

func (c *Consul) Register() error {

	data, err := json.Marshal(c.registerMeta)
	if err != nil {
		return fmt.Errorf("Consul.Register json.Marshal: %v", err)
	}

	log.Debugf("data: %s", string(data))

	endpoint := fmt.Sprintf("%s/v1/catalog/register", c.endpoint)

	req, err := http.NewRequest("PUT", endpoint, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("Consul.Register http.NewRequest: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("Consul.Register c.client.Do: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Consul.Register: failed to register client")
	}

	return nil
}
