package main

import (
	ctrl "github.com/davidesalerno/workforce-shift-scheduling/pkg/controller"
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/schedule", ctrl.RestScheduleHandler)
	router.Use(cors.Default())
	router.Run(":8080")
}
