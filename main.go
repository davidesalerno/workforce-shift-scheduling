package main

import (
	"time"

	ctrl "github.com/davidesalerno/workforce-shift-scheduling/pkg/controller"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router := gin.Default()
	router.Use(cors.New(config))
	router.POST("/schedule", ctrl.RestScheduleHandler)
	router.Run(":8080")
}
