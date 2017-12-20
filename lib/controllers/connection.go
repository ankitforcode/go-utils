package controllers

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"

	log "github.com/inconshreveable/log15"

	"github.com/ankitforcode/go-utils/config"
)

var (
	mysqlConnMaster, mysqlConnSlave, mysqlConn *sql.DB
)

func MySQLConn() (*sql.DB, error) {
	var err error
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQL.User, config.Config.MySQL.Password, config.Config.MySQL.Host, config.Config.MySQL.Port)
	mysqlConn, err = sql.Open("mysql", connection)
	err = mysqlConn.Ping()
	if err != nil {
		log.Error("Failed Connection to Master MySQL Server: ", "err", err.Error())
		return nil, err
	}
	log.Info("Successfully Connected to Master MySQL Server: ", "server", config.Config.MySQL.Host)
	mysqlConn.SetConnMaxLifetime(time.Minute * config.Config.MySQL.ConnLifeTime)
	mysqlConn.SetMaxIdleConns(config.Config.MySQL.MaxIdle)
	mysqlConn.SetMaxOpenConns(config.Config.MySQL.MaxOpen)
	return mysqlConn, nil
}

func MySQLMasterConn() (*sql.DB, error) {
	var err error
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLMaster.User, config.Config.MySQLMaster.Password, config.Config.MySQLMaster.Host, config.Config.MySQLMaster.Port)
	mysqlConnMaster, err = sql.Open("mysql", connection)
	err = mysqlConnMaster.Ping()
	if err != nil {
		log.Error("Failed Connection to Master MySQL Server: ", "err", err.Error())
		return nil, err
	}
	log.Info("Successfully Connected to Master MySQL Server: ", "server", config.Config.MySQLMaster.Host)
	mysqlConnMaster.SetConnMaxLifetime(time.Minute * config.Config.MySQLMaster.ConnLifeTime)
	mysqlConnMaster.SetMaxIdleConns(config.Config.MySQLMaster.MaxIdle)
	mysqlConnMaster.SetMaxOpenConns(config.Config.MySQLMaster.MaxOpen)
	return mysqlConnMaster, nil
}

func MySQLSlaveConn() (*sql.DB, error) {
	var err error
	var connection = fmt.Sprintf("%s:%s@tcp(%s:%d)/", config.Config.MySQLSlave.User, config.Config.MySQLSlave.Password, config.Config.MySQLSlave.Host, config.Config.MySQLSlave.Port)
	mysqlConnSlave, err = sql.Open("mysql", connection)
	err = mysqlConnSlave.Ping()
	if err != nil {
		log.Error("Failed Connection to Slave MySQL Server: ", "err", err.Error())
		return nil, err
	}
	log.Info("Successfully Connected to Slave MySQL Server: ", "server", config.Config.MySQLSlave.Host)
	mysqlConnSlave.SetConnMaxLifetime(time.Minute * config.Config.MySQLSlave.ConnLifeTime)
	mysqlConnSlave.SetMaxIdleConns(config.Config.MySQLSlave.MaxIdle)
	mysqlConnSlave.SetMaxOpenConns(config.Config.MySQLSlave.MaxOpen)
	return mysqlConnSlave, nil
}
