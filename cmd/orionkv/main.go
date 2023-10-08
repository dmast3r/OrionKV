package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dmazt3r/OrionKV/internal/config"
	"github.com/dmazt3r/OrionKV/internal/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", handlers.GetHandler).Methods("GET")
	router.HandleFunc("/put", handlers.PutHandler).Methods("POST")
	router.HandleFunc("/delete", handlers.DeleteHandler).Methods("DELETE")

	fmt.Printf("Starting Orion on the port: %s", config.PORT)

	err := http.ListenAndServe(":" + config.PORT, router)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
		os.Exit(1)
	}
}