package handler

import (
	"net/http"
	"log"
	"encoding/json"
	"github.com/TechBowl-japan/go-stations/model"
)

// A HealthzHandler implements health check endpoint.
type HealthzHandler struct{}

// NewHealthzHandler returns HealthzHandler based http.Handler.
func NewHealthzHandler() *HealthzHandler {
	return &HealthzHandler{}
}

// ServeHTTP implements http.Handler interface.
func (h *HealthzHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := &model.HealthzResponse{"OK"}
	response,err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}
	text := string(response) + "\n"
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(text))
}
