package config

import (
	"fmt"

	"github.com/r3boot/rkt-buddy/lib/system"
)

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	sysdata, err := system.GetSystemData()
	if err != nil {
		return nil, fmt.Errorf("LoadConfig: %v", err)
	}

	if cfg.Service.Node == "" && sysdata.Hostname != "" {
		cfg.Service.Node = sysdata.Hostname
	}

	if cfg.Service.Address == "" && sysdata.Address != "" {
		cfg.Service.Address = sysdata.Address
	}

	if cfg.Service.Interface == "" && sysdata.Interface != "" {
		cfg.Service.Interface = sysdata.Interface
	}

	return cfg, nil
}
