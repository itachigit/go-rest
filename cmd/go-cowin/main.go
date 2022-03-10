package main

import (
	"net/http"

	"go-cowin/api"
	"go-cowin/config"
	"go-cowin/db/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	cfg := config.LoadConfig()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	db := cfg.Connect()
	db.AutoMigrate(&models.State{})
	cowinServer := api.NewCowinServer(cfg, logger, db)
	r.HandleFunc("/", cowinServer.GetState)
	logger.Info("Starting the server")
	svr := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	svr.ListenAndServe()
}
