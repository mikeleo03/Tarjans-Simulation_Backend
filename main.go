package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mikeleo03/Tarjans-Algorithm_Backend/router"
)

func main() {
	// Activate router
	r := router.Router()

    // Start server
	// Use `PORT` provided in environment or default to 3000
  	var port = envPortOr("8080")
	fmt.Println("Starting server...")
	fmt.Println("Listening from" + port)
	log.Fatal(http.ListenAndServe(port, r))
	fmt.Println("The server is started!")
}

// PORT handling
// Returns PORT from environment if found, defaults to
// value in `port` parameter otherwise. The returned port
// is prefixed with a `:`, e.g. `":3000"`.
func envPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}