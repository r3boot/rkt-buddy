package api

import "github.com/r3boot/rkt-buddy/lib/logger"

var (
	log *logger.Logger
)

func NewApi(l *logger.Logger, address string, port int) *Api {
	log = l

	return &Api{
		Address:      address,
		Port:         port,
		HealthEvents: make(chan bool),
	}
}
