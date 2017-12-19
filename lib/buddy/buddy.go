package buddy

import (
	"fmt"

	"github.com/r3boot/rkt-buddy/lib/api"
	"github.com/r3boot/rkt-buddy/lib/config"
	"github.com/r3boot/rkt-buddy/lib/consul"
	"github.com/r3boot/rkt-buddy/lib/logger"
)

var (
	log *logger.Logger
)

func NewBuddy(l *logger.Logger, cfg *BuddyConfig) (*Buddy, error) {
	log = l

	config, err := config.LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("NewBuddy: %v", err)
	}

	consul, err := consul.NewConsul(log, config)
	if err != nil {
		return nil, fmt.Errorf("NewBuddy: %v", err)
	}

	b := &Buddy{
		cfg:          config,
		consul:       consul,
		api:          api.NewApi(log, config.Service.Address, 6548),
		healthEvents: make(chan bool),
	}

	if cfg.HealthCheck != nil {
		b.healthCheck = cfg.HealthCheck
	} else {
		b.healthCheck = DummyHealthCheck
	}

	return b, nil
}
