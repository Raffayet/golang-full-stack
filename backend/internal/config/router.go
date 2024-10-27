package config

import (
	"data-merging/internal/api"

	"github.com/gorilla/mux"
)

func SetupRouter(profileHandler *api.DatasetHandler) *mux.Router {
	// Setup HTTP routing
	router := mux.NewRouter()
	router.HandleFunc("/datasets", profileHandler.GetDatasetHandler).Methods("GET")
	return router
}
