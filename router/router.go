package router

import (
	"receipt-processor-challenge/controllers"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.POST("/receipts/process", controllers.PostReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)

	router.Run(":8080")
}