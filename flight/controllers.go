package flight

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddFlightsInput []string

type Flights struct {
	From string `json:"from"`
	To   string `json:"to`
}

type GetFlightsResponse []Flights

func AddFlights(c *gin.Context) {
	var requestBody AddFlightsInput
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{"error": err})
		return
	}
	err := AddFlightModel(requestBody[0], requestBody[1])
	if err != nil {
		if err.Error() == "Duplicate Error" {
			c.AbortWithStatus(http.StatusConflict)
			return
		}
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}

func GetFlights(c *gin.Context) {
	flights, err := GetFlightsModel()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var getFlightsResponse GetFlightsResponse
	for _, f := range flights {
		getFlightsResponse = append(getFlightsResponse, Flights{From: f.From, To: f.To})

	}
	c.JSON(http.StatusOK, getFlightsResponse)
}

func GetFlightsFrom(c *gin.Context) {
	log.Println("here")
	from := c.Param("from")
	flights, err := GetFlightsFromModel(from)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	var getFlightsResponse GetFlightsResponse
	for _, f := range flights {
		getFlightsResponse = append(getFlightsResponse, Flights{From: f.From, To: f.To})

	}
	c.JSON(http.StatusOK, getFlightsResponse)
}
