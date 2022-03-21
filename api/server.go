package api

import (
	"go-rest/config"

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

func (svr *CowinServer) GetLogger(val string) *logrus.Entry {
	logger := svr.logger.WithFields(logrus.Fields{
		"method": val,
	})
	return logger
}
