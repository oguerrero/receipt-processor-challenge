package controllers

import (
	"net/http"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var processedReceipts []models.ProcessedReceipt

func PostReceipt(c *gin.Context) {
	var receipt models.Receipt
	err := c.BindJSON(&receipt)

	if err != nil {
		c.String(http.StatusBadRequest, "The receipt is invalid")
		return
	}

	points := utils.CalculatePoints(receipt)

	id := uuid.New().String()
	processedReceipts = append(processedReceipts, models.ProcessedReceipt{ID: id, Points: points})

	c.String(http.StatusOK, id)
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")

	for _, processedReceipt := range processedReceipts {
		if processedReceipt.ID == id {
			c.JSON(http.StatusOK, processedReceipt.Points)
			return
		}
	}
	
	c.String(http.StatusNotFound, "No receipt found for that id")
}