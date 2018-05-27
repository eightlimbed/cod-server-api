package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	for {
		var router = mux.NewRouter()
		router.HandleFunc("/status", healthStatus).Methods("GET")
		router.HandleFunc("/start", handleStart).Methods("GET")

		// For CORS
		headersOk := handlers.AllowedHeaders([]string{"Authorization"})
		originsOk := handlers.AllowedOrigins([]string{"*"})
		methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

		fmt.Println("Server running on port 3000...")
		log.Fatal(http.ListenAndServe(":3000",
			handlers.CORS(originsOk, headersOk, methodsOk)(router)))
	}
}

func handleStart(w http.ResponseWriter, r *http.Request) {
	cmdString := []string{"docker", "run", "-itd", "ubuntu:14.04"}
	cmd := exec.Command(cmdString)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
	fmt.Println("Command executed.")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("OK")
}
