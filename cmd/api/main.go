package main

import (
	"StudentStats-backend-go/server/config"
	"StudentStats-backend-go/server/routes"
	"StudentStats-backend-go/server/ws"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB

func main() {
	
	DB = config.InitDB()

	r := gin.Default()
	
	websocketHub := ws.NewHub()
	go websocketHub.Run()

	routes.Setup(r, websocketHub, DB)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
