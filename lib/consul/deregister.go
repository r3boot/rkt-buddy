package consul

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type DeregisterData struct {
	Datacenter string `json:"Datacenter"`
	Node       string `json:"Node"`
}

func (c *Consul) Deregister() error {
	data, err := json.Marshal(c.deregisterCatalogMeta)
	if err != nil {
		return fmt.Errorf("Consul.Register json.Marshal: %v", err)
	}

	endpoint := fmt.Sprintf("%s/v1/catalog/deregister", c.endpoint)

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
