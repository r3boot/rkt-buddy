package api

import "net/http"

func (a *Api) HealthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			msg := "Service is DOWN"
			if a.serviceIsUp {
				msg = "Service is up"
			}

			response := HealthResponse{
				Status:  a.serviceIsUp,
				Message: msg,
			}

			data, err := response.ToJSON()
			if err != nil {
				log.Warningf("Api.HealthHandler: %v", err)
				errorResponse(w, r, "json encoding failed")
				return
			}

			w.Header().Add("Content-Type", "application/json")
			w.Write(data)
			okResponse(r, len(data))
		}
	default:
		{
			errorResponse(w, r, "Unsupported method")
		}
	}
}
