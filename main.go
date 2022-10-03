package main

import (
	"github.com/hugosjoberg/flight_tracker/config"
	"github.com/hugosjoberg/flight_tracker/db"
	"github.com/hugosjoberg/flight_tracker/migrate"
	"github.com/hugosjoberg/flight_tracker/server"
)

func main() {
	config.Init()
	db.Init()
	migrate.Migrate()
	server.Init()
}
