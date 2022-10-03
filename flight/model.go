package flight

import (
	"errors"
	"log"

	"github.com/google/uuid"
	"github.com/hugosjoberg/flight_tracker/db"
	"gorm.io/gorm"
)

type Flight struct {
	gorm.Model
	ID   string `gorm:"primaryKey"`
	From string
	To   string
}

func AddFlightModel(from string, to string) error {
	db := db.DB

	// Check that entry does not exist
	var exists bool
	if err := db.Model(&Flight{}).Select("count(*) > 0").Where(`flights.from = ? AND flights.to = ?`, from, to).Group("id").First(&exists).Error; err != nil {
		log.Println(err)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	if exists {
		return errors.New("Duplicate Error")
	}

	id := uuid.New()
	flight := &Flight{ID: id.String(), From: from, To: to}
	if err := db.Create(flight).Error; err != nil {
		log.Println("error - unable to create record")
		return err
	}
	return nil
}

func FindFlightModel(from string, to string) (*Flight, error) {
	db := db.DB

	var flight Flight
	if err := db.Where(&Flight{From: from, To: to}).First(&flight).Error; err != nil {
		log.Println("error - unable to query flight")
		return nil, err
	}
	return &flight, nil
}

func GetFlightsModel() ([]Flight, error) {
	db := db.DB
	flights := []Flight{}
	if err := db.Find(&flights).Error; err != nil {
		log.Println("error - unable to query flight")
		return nil, err
	}
	return flights, nil

}

func GetFlightsFromModel(from string) ([]Flight, error) {
	db := db.DB
	flights := []Flight{}
	if err := db.Where(&Flight{From: from}).Find(&flights).Error; err != nil {
		log.Println("error - unable to query flight")
		return nil, err
	}
	return flights, nil

}
