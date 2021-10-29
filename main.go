package main

import (
	"log"
	"mnc-backend/cfg"
	"mnc-backend/http"
	"mnc-backend/repositories"

	_ "github.com/lib/pq"
)

func init() {
	cfg.Init()
}

func main() {
	app := http.App{}

	conn, err := repositories.Conn()
	if err != nil {
		log.Printf("Cannot initialize connection to database: %v", err)
	}

	app.Initialize(conn)
	app.RunServer()
}
