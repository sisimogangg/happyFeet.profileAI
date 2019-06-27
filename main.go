package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	controller "happyFeet.profileAI/controllers"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/api/profile/{userId}", controller.GetProfile).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Println(port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}

}
