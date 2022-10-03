package migrate

import (
	"errors"
	"fmt"
	"log"

	"github.com/hugosjoberg/flight_tracker/db"
	"github.com/hugosjoberg/flight_tracker/flight"
	"gorm.io/gorm"
)

func Migrate() error {
	log.Println("Running migrations")
	dbConn := db.DB
	dbConn.AutoMigrate(&flight.Flight{})

	// Let's seed the DB with some dummy data if it doesn't exist
	_, err := flight.FindFlightModel("SFO", "EWR")
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil
	}
	log.Println("seeding database")
	// The database is not seeded, let's add some data
	flights := [...][2]string{
		{"SFO", "EWR"}, {"ATL", "EWR"}, {"IND", "EWR"},
		{"SFO", "ATL"}, {"GSO", "IND"}, {"ATL", "GSO"},
	}
	for _, f := range flights {
		err = flight.AddFlightModel(f[0], f[1])
		if err != nil {
			fmt.Println("flight ", f[0], f[1], " already exists")
		}
	}
	return nil
}
