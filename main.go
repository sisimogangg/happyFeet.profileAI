package main


import (
	"github.com/gorilla/mux"
	"os"
	"fmt"
	"net/http"
)

func main () {
   router := mux.NewRouter()


   port := os.Getenv("PORT")
   if port == "" {
	   port = "8000"
   }

   fmt.Println(port)

   err := http.ListenAndServe(":" + port, router)
   if err != nil {
	   fmt.Print(err)
   }

}