package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/itachigit/goREST/api"
	"github.com/itachigit/goREST/config"
	"github.com/itachigit/goREST/db/models"
	"github.com/sirupsen/logrus"
)

/*
func handler(w http.ResponseWriter, r *http.Request) { // this is it
	w.Write([]byte("Hello!\n"))
} */

func main() {
	r := mux.NewRouter()
	config := config.LoadConfig()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	db := config.Connect()
	db.AutoMigrate(&models.State{})
	cowinServer := api.NewCowinServer(config, logger, db)
	r.HandleFunc("/", cowinServer.GetState)
	log.Println("Starting the server")
	svr := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	svr.ListenAndServe()
}
