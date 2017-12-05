package config

import (
	"os"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/jinzhu/configor"
)

// Config Default
var Config = struct {
	Server struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"7000" env:"LISTEN_PORT"`
		Loglevel string `default:"info" json:"loglevel"`
	}
	MySQLMaster struct {
		User         string        `default:"root"`
		Password     string        `default:"root"`
		Host         string        `default:"localhost"`
		Port         uint          `default:"3306"`
		MaxIdle      int           `default:"0"`
		MaxOpen      int           `default:"10"`
		ConnLifeTime time.Duration `default:"5"`
	}
	MySQLSlave struct {
		User         string        `default:"root"`
		Password     string        `default:"root"`
		Host         string        `default:"localhost"`
		Port         uint          `default:"3306"`
		MaxIdle      int           `default:"0"`
		MaxOpen      int           `default:"10"`
		ConnLifeTime time.Duration `default:"5"`
	}
	Redis struct {
		Host      string `default:"localhost"`
		Port      uint   `default:"6379"`
		MaxIdle   int    `default:"80"`
		MaxActive int    `default:"12000"`
	}
}{}

// Load Configurations
func init() {
	env := os.Getenv("GOLANG_ENV")
	os.Setenv("CONFIGOR_ENV", os.Getenv("GOLANG_ENV"))
	if env == "" {
		env = "dev"
		os.Setenv("CONFIGOR_ENV", env)
	}
	log.Info("Loading Configuration For :", "env", env)
	if err := configor.Load(&Config, "config/config.json"); err != nil {
		log.Error(err.Error())
	}
}
