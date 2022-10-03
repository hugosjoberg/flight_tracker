package db

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/hugosjoberg/flight_tracker/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	config := config.GetConfig()
	dsn := url.URL{
		User:   url.UserPassword(config.DBUser, config.DBPassword),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", config.DBPHOST, config.DBPort),
		Path:   config.Database,
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn.String(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{PrepareStmt: true})
	if err != nil {
		log.Fatal(err)
	}

	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Could not create database connection: ", err)
	}

	sqlDB.Ping()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	DB = db
}
