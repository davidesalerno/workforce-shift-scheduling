package main

import (
	ctrl "github.com/davidesalerno/workforce-shift-scheduling/pkg/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/schedule", ctrl.RestScheduleHandler)
	router.Run(":8080")
}
