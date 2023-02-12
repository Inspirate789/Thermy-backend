package main

import (
	"logger/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

var log logger.Logger = logger.NewInfluxLogger()

func main() {
	err := log.Open()
	if err != nil {
		panic(err)
	}
	defer log.Close()

	router := gin.Default()
	// router.SetTrustedProxies([]string{"192.168.52.38"})

	router.POST("/log", postMessage)

	router.Run("0.0.0.0:8080")
}

func postMessage(c *gin.Context) {
	var r logger.LogRecord

	// Call BindJSON to bind the received JSON to LogRecord
	err := c.BindJSON(&r)
	if err != nil {
		return
	}

	err = log.Print(r)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(http.StatusCreated, r)
}
