package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"StudentStats-backend-go/server/routes"
	"StudentStats-backend-go/server/config"
)

func main() {

	config.InitDB()

	r := gin.Default()

	routes.Setup(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
