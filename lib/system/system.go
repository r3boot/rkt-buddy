package system

import (
	"fmt"
	"os"
)

func GetSystemData() (*SystemData, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("GetSystemData os.Hostname: %v", err)
	}

	intf, err := GetFirstInterface()
	if err != nil {
		return nil, fmt.Errorf("GetSystemData: %v", err)
	}

	address, err := GetAddress(intf)
	if err != nil {
		return nil, fmt.Errorf("GetSystemData: %v", err)
	}

	return &SystemData{
		Hostname:  hostname,
		Interface: intf,
		Address:   address,
	}, nil
}
