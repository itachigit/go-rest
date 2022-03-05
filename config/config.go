package config

import (
	"fmt"

	envstruct "code.cloudfoundry.org/go-envstruct"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Cowin struct {
	Url string `env:"COWIN_URL, required, report"`
}

type DB struct {
	DbName string `env:"DB_NAME, required, report"`
	DbUser string `env:"DB_USER, required, report"`
	DbPass string `env:"DB_PASS, required"`
	DbHost string `env:"DB_HOST, required, report"`
	DbPort string `env:"DB_PORT, required, report"`
}

type Config struct {
	Cowin Cowin
	DB    DB
}

func LoadConfig() *Config {
	config := Config{}
	err := envstruct.Load(&config)
	if err != nil {
		log.Error("Error loading environment variables..")
	}
	err = envstruct.WriteReport(&config)
	if err != nil {
		log.Error("Error writing report..")
	}
	return &config
}

func (cfg *Config) Connect() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB.DbHost, cfg.DB.DbUser, cfg.DB.DbPass, cfg.DB.DbName, cfg.DB.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("error-connecting-database", err)
	}
	return db
}
