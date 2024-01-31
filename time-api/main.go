package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"time"
	"strings"
)

func getTime(c *gin.Context) {
	response := make(map[string]string)
	timezones := strings.Split(c.Query("tz"), ",")

	if timezones[0] == "" {
		loc, _ := time.LoadLocation("UTC")
		time := time.Now().In(loc).String()
		response["current_time"] = time
		c.JSON(http.StatusOK, response)

		return
	}

	for i, zone := range timezones {
		loc, err := time.LoadLocation(zone)

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "invalid timezone",
			})

			return
		}

		time := time.Now().In(loc).String()
		fmt.Println(i, zone, time)
		response[zone] = time
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()
	router.GET("/api/time", getTime)

	router.Run("localhost:3000")
}