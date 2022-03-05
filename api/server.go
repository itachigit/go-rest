package api

import (
	"github.com/itachigit/goREST/config"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type CowinServer struct {
	config *config.Config
	logger *logrus.Logger
	db     *gorm.DB
}

func NewCowinServer(config *config.Config, logger *logrus.Logger, db *gorm.DB) CowinServer {
	return CowinServer{
		config: config,
		logger: logger,
		db:     db,
	}
}
