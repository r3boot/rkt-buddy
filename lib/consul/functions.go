package consul

func (c *Consul) RunAgent() {
	c.Deregister()

	for serviceIsUp := range c.HealthEvents {
		if serviceIsUp {
			log.Debugf("Consul.RunAgent: registering service with consul")
			c.Register()
		} else {
			log.Debugf("Consul.RunAgent: deregistering service with consul")
			c.Deregister()
		}
	}
}
