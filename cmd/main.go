package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/denizedizcan/Golang-Curmin/common/config/envs"
	"github.com/denizedizcan/Golang-Curmin/common/db"
	"github.com/denizedizcan/Golang-Curmin/handlers"
	"github.com/denizedizcan/Golang-Curmin/middlewares"
	"github.com/gorilla/mux"
)

// start the app and handle routes
func main() {
	fmt.Println("Starting App..")
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}
	DB := db.Init(c.User, c.Password, c.Host, c.Port, c.DBname)

	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("api/currencies", middlewares.SetMiddlewareJSON(h.ShowCurrency)).Methods("POST")
	/*
		router.HandleFunc("api/timeseries", middlewares.SetMiddlewareJSON()).Methods("GET")
		router.HandleFunc("api/latest", middlewares.SetMiddlewareJSON()).Methods("GET")
	*/
	http.ListenAndServe(":8080", router)
}
