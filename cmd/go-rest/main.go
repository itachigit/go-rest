package main

import (
	"net/http"

	"go-rest/api"
	"go-rest/config"
	"go-rest/db/models"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {
	r := mux.NewRouter()
	cfg := config.LoadConfig()
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	db, err := cfg.Connect()
	if err != nil {
		logger.Error("error-connecting-database: ", err)
		return
	}
	err = db.AutoMigrate(&models.Entry{})
	if err != nil {
		logger.Error("Error while migrating database: ", err)
		return
	}
	cowinServer := api.NewCowinServer(cfg, logger, db)
	r.HandleFunc("/", cowinServer.GetState)
	logger.Info("Starting the server")
	svr := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}
	svr.ListenAndServe()
}
