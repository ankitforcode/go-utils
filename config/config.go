package config

import (
	"os"
	"time"

	log "github.com/inconshreveable/log15"
	"github.com/jinzhu/configor"
)

// Config Default
var Config = struct {
	Entity string
	Name   string
	Server struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"7000" env:"LISTEN_PORT"`
		Loglevel string `default:"info" json:"loglevel"`
	}
	AWS struct {
		AccessKey         string `json:"aws_access_key"`
		SecretKey         string `json:"aws_secret_key"`
		ImageUploadFolder string `json:"image_upload_folder"`
		DstImageFolder    string `json:"dst_image_folder"`
		Region            string `json:"aws_region"`
	}
	MySQLMaster struct {
		User         string        `default:"root"`
		Password     string        `default:""`
		Host         string        `default:"localhost"`
		Port         uint          `default:"3306"`
		MaxIdle      int           `default:"0"`
		MaxOpen      int           `default:"10"`
		ConnLifeTime time.Duration `default:"5"`
	} `json:"mysqlMaster"`
	MySQLSlave struct {
		User         string        `default:"root"`
		Password     string        `default:""`
		Host         string        `default:"localhost"`
		Port         uint          `default:"3306"`
		MaxIdle      int           `default:"0"`
		MaxOpen      int           `default:"10"`
		ConnLifeTime time.Duration `default:"5"`
	} `json:"mysqlSlave"`
	MySQL struct {
		User         string        `default:"root"`
		Password     string        `default:""`
		Host         string        `default:"localhost"`
		Port         uint          `default:"3306"`
		MaxIdle      int           `default:"0"`
		MaxOpen      int           `default:"10"`
		ConnLifeTime time.Duration `default:"5"`
	} `json:"mysql"`
	Redis struct {
		Host      string `default:"localhost"`
		Port      uint   `default:"6379"`
		MaxIdle   int    `default:"80"`
		MaxActive int    `default:"12000"`
	}
	MongoDB struct {
		User       string
		Password   string
		Hosts      []interface{}
		SSLEnabled bool `json:"sslEnabled"`
	}
	Rabbit struct {
		Host     string `default:"localhost"`
		Port     uint   `default:"5672"`
		User     string `default:"guest"`
		Password string `default:"guest"`
		Vhost    string `default:""`
		Q        map[string]string
	}
	ImageType struct {
		CDN   map[string]interface{} `json:"cdn"`
		Sizes []map[string]interface{}
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
