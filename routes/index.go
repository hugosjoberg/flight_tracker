package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosjoberg/flight_tracker/flight"
)

func AddRoutes(superRoute *gin.RouterGroup) {
	flight.FlightRouter(superRoute)
}
