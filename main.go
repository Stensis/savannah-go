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

    // Log the AUTH0_DOMAIN to verify it's set correctly
    auth0Domain := os.Getenv("AUTH0_DOMAIN")
    log.Printf("AUTH0_DOMAIN: %s", auth0Domain)  // Log the issuer
	log.Printf("AUTH0_DOMAIN: %s", auth0Domain)

    // Create a new router
    r := router.Router()

    // Get port
    Port := os.Getenv("PORT")
    if Port == "" {
        Port = "8080"
    }

    // Set CORS
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

    // Establish logger
    n := negroni.Classic()
    n.UseHandler(r)

    // Establish server
    server := &http.Server{
        Handler: handlers.CORS(originsOk, headersOk, methodsOk)(n),
        Addr:    ":" + Port,
    }

    log.Printf("Listening on PORT: %s", Port)
    server.ListenAndServe()
}
