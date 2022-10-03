package flight

import "github.com/gin-gonic/gin"

func FlightRouter(superRoute *gin.RouterGroup) {
	flightsRouter := superRoute.Group("/flight")
	{
		flightsRouter.POST("/", AddFlights)
		flightsRouter.GET("/", GetFlights)
		flightsRouter.GET("/:from", GetFlightsFrom)

	}
}
