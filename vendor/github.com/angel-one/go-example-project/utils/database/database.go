package database

import (
	"database/sql"
	"fmt"
	"github.com/angel-one/go-example-project/constants"
	"github.com/angel-one/go-utils/log"
	"time"
)

type Config struct {
	Server                string        `json:"server"`
	Port                  int           `json:"port"`
	Name                  string        `json:"name"`
	Username              string        `json:"-"`
	Password              string        `json:"-"`
	MaxOpenConnections    int           `json:"maxOpenConnections"`
	MaxIdleConnections    int           `json:"maxIdleConnections"`
	ConnectionMaxLifetime time.Duration `json:"connectionMaxLifetime"`
	ConnectionMaxIdleTime time.Duration `json:"connectionMaxIdleTime"`
}

var db *sql.DB

func InitDatabase(config Config) error {
	log.Info(nil).Interface(constants.DatabaseConfigKey, config).Msg("initializing database")
	var err error

	// open the database
	db, err = sql.Open(
		constants.MySQLDriverName,
		fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password,
			config.Server, config.Port, config.Name),
	)
	if err != nil {
		return err
	}

	// try to ping
	err = db.Ping()
	if err != nil {
		return err
	}

	// now set the configurations
	db.SetMaxOpenConns(config.MaxOpenConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxIdleTime(config.ConnectionMaxIdleTime)
	db.SetConnMaxLifetime(config.ConnectionMaxLifetime)

	return nil
}

func Get() *sql.DB {
	return db
}

func Close() error {
	return db.Close()
}
