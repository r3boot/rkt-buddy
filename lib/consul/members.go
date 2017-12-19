package consul

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type MemberData struct {
	Name        string            `json:"Name"`
	Addr        string            `json:"Addr"`
	Port        int               `json:"Port"`
	Tags        map[string]string `json:"Tags"`
	Status      int               `json:"Status"`
	ProtocolMin int               `json:"ProtocolMin"`
	ProtocolMax int               `json:"ProtocolMax"`
	ProtocolCur int               `json:"ProtocolCur"`
	DelegateMin int               `json:"DelegateMin"`
	DelegateMax int               `json:"DelegateMax"`
	DelegateCur int               `json:"DelegateCur"`
}

type MembersData []MemberData

func (md *MembersData) ToJSON() ([]byte, error) {
	data, err := json.Marshal(md)
	return data, err
}

func (md *MembersData) FromJSON(data []byte) error {
	if len(data) == 0 {
		return fmt.Errorf("MembersData.FromJSON: got zero length input")
	}
	err := json.Unmarshal(data, md)
	return err
}

func (c *Consul) GetMembers() (*MembersData, error) {
	members := &MembersData{}

	endpoint := c.endpoint + "/v1/agent/members"

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return nil, fmt.Errorf("Consul.GetMembers http.NewRequest: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Consul.GetMembers c.client.Do: %v", err)
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(members)
	if err != nil {
		return nil, fmt.Errorf("Consul.GetMembers decoder.Decode: %v", err)
	}

	return members, nil
}
