package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/viper"

	"github.com/gorilla/mux"
	controller "github.com/sisimogangg/happyFeet.profileAI/controllers"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool("debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
}

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/api/profile/{userId}", controller.GetProfile).Methods("GET")
	router.HandleFunc("/api/profile/{userId}", controller.CreateProfile).Methods("POST")
	//router.Handle("/api/profile/{userId}", controller.GetProfile).

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
