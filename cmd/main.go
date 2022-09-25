package main

import (
	"fmt"
	"log"
	"net/http"

	config "github.com/denizedizcan/Golang-Curmin/common/config/envs"
	"github.com/denizedizcan/Golang-Curmin/common/db"
	"github.com/denizedizcan/Golang-Curmin/common/handlers"
	"github.com/denizedizcan/Golang-Curmin/common/middlewares"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var (
	c   config.Config
	DB  *gorm.DB
	err error
)

// start the app and handle routes
func init() {
	initConfig()
	initDB(c)
}

func main() {
	fmt.Println("Starting App..")

	h := handlers.New(DB)
	router := mux.NewRouter()
	router.HandleFunc("/api/currencies", middlewares.SetMiddlewareJSON(h.ShowCurrency)).Methods("GET")
	router.HandleFunc("/api/currencies/list", middlewares.SetMiddlewareJSON(h.InsertCurrency)).Methods("POST")
	router.HandleFunc("/api/currencies/data/list", middlewares.SetMiddlewareJSON(h.InsertCurrencyData)).Methods("POST")
	router.HandleFunc("/api/currencies/data/all", middlewares.SetMiddlewareJSON(h.InsertAllCurrencyData)).Methods("POST")
	router.HandleFunc("/api/currency/latest", middlewares.SetMiddlewareJSON(h.GetCurrency)).Methods("GET")

	/*
		router.HandleFunc("api/timeseries", middlewares.SetMiddlewareJSON()).Methods("GET")
		router.HandleFunc("api/latest", middlewares.SetMiddlewareJSON()).Methods("GET")
	*/
	http.ListenAndServe(":8080", router)

}

func initDB(c config.Config) {
	DB = db.Init(c.User, c.Password, c.Host, c.Port, c.DBname)
}

func initConfig() {
	c, err = config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}
}
