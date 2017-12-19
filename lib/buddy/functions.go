package buddy

import (
	"time"
)

func (b *Buddy) Run() {
	go b.consul.RunAgent()
	go b.api.Run()
	go b.RunHealthCheck()

	for healthEvent := range b.healthEvents {
		b.api.HealthEvents <- healthEvent
		b.consul.HealthEvents <- healthEvent
	}
}

func (b *Buddy) ServiceIsFunctional() bool {
	return b.serviceIsUp
}

func (b *Buddy) RunHealthCheck() {
	numChecks := 0
	serviceIsUp := false

	for {
		serviceIsHealty := b.healthCheck()

		if serviceIsHealty != serviceIsUp {
			if numChecks >= NUM_CONSECUTIVE_CHECKS {
				b.healthEvents <- serviceIsHealty
				serviceIsUp = serviceIsHealty
				numChecks = 0
			} else {
				numChecks += 1
			}
		}

		time.Sleep(1 * time.Second)
	}
}
