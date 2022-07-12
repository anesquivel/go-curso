package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HelloMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Hello")
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func ByeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Bye")
		// Before calling handler
		c.Next()
		// After calling handler
	}
}

func main() {
	router := gin.Default()
	router.Use(HelloMiddleware(), ByeMiddleware())
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Middleware works!"})
	})
	router.Run(":8080")
}
