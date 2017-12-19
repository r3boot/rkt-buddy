package consul

import (
	"net/http"

	"github.com/r3boot/rkt-buddy/lib/config"
)

type Consul struct {
	config         *config.Config
	endpoint       string
	client         *http.Client
	HealthEvents   chan bool
	registerMeta   RegisterData
	deregisterMeta DeregisterData
}
