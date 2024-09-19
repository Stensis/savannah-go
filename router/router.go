package router

import (
	"go-service/authenticator"
	"go-service/lib"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)

	router.HandleFunc("/callback", lib.Callback).Methods("POST", "OPTIONS")

	// Protected routes
	router = router.PathPrefix("/api/v1").Subrouter()

	// Protect routes with middleware
	router.Use(authenticator.AuthMiddleware)
	router.HandleFunc("/customers", lib.CreateCustomer).Methods("POST", "OPTIONS")
	router.HandleFunc("/orders", lib.CreateOrder).Methods("POST", "OPTIONS")

	return router
}
