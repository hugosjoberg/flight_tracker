package server

import (
	"github.com/gin-gonic/gin"
	"github.com/hugosjoberg/flight_tracker/routes"
)

func NewRouter() *gin.Engine {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())

	router := app.Group("/api/v1")
	routes.AddRoutes(router)
	return app

}
