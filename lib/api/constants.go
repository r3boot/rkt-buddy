package api

const (
	TF_CLF = "02/Jan/2006:15:04:05 -0700"
)

type HealthResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type Api struct {
	Address      string
	Port         int
	serviceIsUp  bool
	HealthEvents chan bool
}
