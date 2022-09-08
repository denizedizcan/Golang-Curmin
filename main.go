package main

import (
	"fmt"
	"net/http"

	"github.com/denizedizcan/Golang-Curmin/db"
	"github.com/denizedizcan/Golang-Curmin/handlers"
	"github.com/denizedizcan/Golang-Curmin/middlewares"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// start the app and handle routes
func main() {
	fmt.Println("Starting App..")
	godotenv.Load()
	DB := db.Init()
	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("api/currencies", middlewares.SetMiddlewareJSON(h.ShowCurrency)).Methods("POST")
	/*
		router.HandleFunc("api/timeseries", middlewares.SetMiddlewareJSON()).Methods("GET")
		router.HandleFunc("api/latest", middlewares.SetMiddlewareJSON()).Methods("GET")
	*/
	http.ListenAndServe(":8080", router)
}
