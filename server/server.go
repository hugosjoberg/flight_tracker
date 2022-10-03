package server

import (
	"fmt"

	"github.com/hugosjoberg/flight_tracker/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%d", config.Port))
}
