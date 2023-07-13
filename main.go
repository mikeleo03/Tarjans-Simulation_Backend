package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mikeleo03/Tarjans-Algorithm_Backend/router"
)

func main() {
	// Activate router
	r := router.Router()

    // Start server
	fmt.Println("Starting server on the port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}