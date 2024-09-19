package main

import (
	"go-service/db"
	"go-service/router"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/urfave/negroni"
)


func main() {
	// Initialize the database
	db.InitDB()

	// Create a new router
	r := router.Router()

	//Get port
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}

	// set CORS
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	// establish logger
	n := negroni.Classic()
	n.UseHandler(r)

	// establish server
	server := &http.Server{
		Handler: handlers.CORS(originsOk, headersOk, methodsOk)(n),
		Addr:    ":" + Port,
	}

	log.Printf("Listening on PORT: %s", Port)
	server.ListenAndServe()
}
