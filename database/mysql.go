package database

import (
	"coroner/config"
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
)

var Db *sql.DB

func Init() {
	logrus.Info("initializing databases' connection")
	if Db != nil {
		return
	}
	db, err := create()
	if err != nil {
		logrus.Error("failed to connect to db ", err)
		return
	}
	logrus.Info("connected successfully")

	Db = db
}

func Close() {
	logrus.Info("closing databases' connection")
	Db.Close()
}



func create() (*sql.DB, error) {
	db, err := sql.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&collation=utf8_unicode_ci&readTimeout=%s",
			config.Cfg.Mysql.Username,
			config.Cfg.Mysql.Password,
			config.Cfg.Mysql.Host,
			config.Cfg.Mysql.Port,
			config.Cfg.Mysql.DBName,
			config.Cfg.Mysql.ConnectTimeout,
		),
	)
	if err != nil {
		log.Panicln(err)
	}

	db.SetMaxOpenConns(config.Cfg.Mysql.MaxOpenConnections)
	db.SetMaxIdleConns(config.Cfg.Mysql.MaxIdleConnections)

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}