package buddy

import (
	"github.com/r3boot/rkt-buddy/lib/api"
	"github.com/r3boot/rkt-buddy/lib/config"
	"github.com/r3boot/rkt-buddy/lib/consul"
)

const (
	NUM_CONSECUTIVE_CHECKS = 3
)

type BuddyConfig struct {
	HealthCheck func() bool
}

type Buddy struct {
	cfg          *config.Config
	consul       *consul.Consul
	api          *api.Api
	serviceIsUp  bool
	healthCheck  func() bool
	healthEvents chan bool
}
