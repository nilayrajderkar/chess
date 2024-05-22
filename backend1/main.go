package main

import (
	routemanager "chess/routemanager"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Chess Lobby")
	routemanager.SetupRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

