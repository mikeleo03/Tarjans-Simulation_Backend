package router

import (
	"github.com/mikeleo03/Tarjans-Algorithm_Backend/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/scc", middleware.ProcessSCC).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/bridges", middleware.ProcessBridges).Methods("POST", "OPTIONS")

	return router
}