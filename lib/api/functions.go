package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (r HealthResponse) ToJSON() ([]byte, error) {
	data, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("WebResponse.ToJSON json.Marshal: %v", err)
	}

	return data, nil
}

func (a *Api) Run() {
	http.HandleFunc("/v1/health", a.HealthHandler)

	go func() {
		for healthEvent := range a.HealthEvents {
			a.serviceIsUp = healthEvent
		}
	}()

	address := fmt.Sprintf("%s:%d", a.Address, a.Port)

	log.Debugf("Api.Run: listening on %s", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatalf("Api.Run http.ListenAndServe: %v", err)
	}
}
