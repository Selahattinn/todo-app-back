package api

import (
	"net/http"

	"github.com/Selahattinn/todo-app-back/pkg/api/response"
	"github.com/Selahattinn/todo-app-back/pkg/service"
	"github.com/gorilla/mux"
)

// Config represents the API configuration
type Config struct {
	Domain        string `yaml:"domain"`
	SigningSecret string `yaml:"signing_secret"`
}

// API represents the structure of the API
type API struct {
	Router  *mux.Router
	config  *Config
	service service.Service
}

// New returns the api settings
func New(config *Config, router *mux.Router, svc service.Service) (*API, error) {
	api := &API{
		config: config,
		//db:     db,
		Router:  router,
		service: svc,
	}

	// Endpoint for browser preflight requests
	api.Router.Methods("OPTIONS").HandlerFunc(api.corsMiddleware(api.preflightHandler))

	// Endpoint for healtcheck
	api.Router.HandleFunc("/api/v1/health", api.corsMiddleware(api.logMiddleware(api.healthHandler))).Methods("GET")

	// Endpoints for jobs
	api.Router.HandleFunc("/api/v1/jobs", api.corsMiddleware(api.logMiddleware(api.GetJobs))).Methods("GET")

	return api, nil
}

func (a *API) healthHandler(w http.ResponseWriter, r *http.Request) {
	response.Write(w, r, struct {
		Status string `json:"status"`
	}{
		"ok",
	})

	return
}

func (a *API) preflightHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	return
}
