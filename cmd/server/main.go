package main

import (
	"log"
	"net/http"
	"os"

	"duty-alert-func/internal/handlers"
)

func main() {
	port := os.Getenv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/CallDutyEngineer", handlers.CallDutyEngineer)

	log.Printf("Listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
