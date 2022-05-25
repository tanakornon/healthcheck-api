package apis

import (
	"backend/src/utils"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Route struct {
	Name     string
	Method   string
	Path     string
	Endpoint func(w http.ResponseWriter, r *http.Request)
}

func NewRouter(container Controllers) http.Handler {
	api := []Route{
		{
			Name:     "Pong",
			Path:     "/",
			Method:   http.MethodGet,
			Endpoint: container.Common.Pong,
		},
		{
			Name:     "Health Check",
			Path:     "/healthcheck",
			Method:   http.MethodPost,
			Endpoint: container.HealthCheck.FindSiteStatus,
		},
	}

	router := mux.NewRouter()
	router.NotFoundHandler = http.HandlerFunc(utils.HandleNotFound)
	router.MethodNotAllowedHandler = http.HandlerFunc(utils.HandleMethodNotAllowed)

	for _, r := range api {
		router.HandleFunc(r.Path, r.Endpoint).Methods(r.Method)
	}

	handler := handlers.CORS()(router)

	return handler
}
