package repositories

import (
	"fmt"
	"log"
	"mnc-backend/cfg"

	"github.com/gocraft/dbr"
	"github.com/spf13/viper"
)

func GetPostgresDSN() (dsn string) {
	dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
		viper.GetString(cfg.PostgresHost),
		viper.GetString(cfg.PostgresUsername),
		viper.GetString(cfg.PostgresPassword),
		viper.GetString(cfg.PostgresDatabaseName))
	log.Printf("%s", dsn)
	return
}

func Conn() (*dbr.Connection, error) {
	conn, err := dbr.Open("postgres", GetPostgresDSN(), nil)

	return conn, err
}
