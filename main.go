package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"./controllers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", controllers.HelloWorld).Methods("GET")
	router.HandleFunc("/api/users", controllers.CreateUser).Methods("POST")
	http.Handle("/", router)

	err := http.ListenAndServe(":8080", router) //Launch the app, visit localhost:8000/api
	if err != nil {
		fmt.Print(err)
	}
}
