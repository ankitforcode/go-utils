package config

import (
	"os"

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
		panic(err)
	}
}
