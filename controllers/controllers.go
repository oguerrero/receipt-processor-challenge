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
		c.JSON(http.StatusBadRequest, gin.H{"Parsing error, please check your request body": err.Error()})
		return
	}

	points := utils.CalculatePoints(receipt)

	id := uuid.New().String()
	processedReceipts = append(processedReceipts, models.ProcessedReceipt{ID: id, Points: points})

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
	id := c.Param("id")

	for _, processedReceipt := range processedReceipts {
		if processedReceipt.ID == id {
			c.JSON(http.StatusOK, gin.H{"points": processedReceipt.Points})
			return
		}
	}
	
	c.JSON(http.StatusNotFound, gin.H{"error": "Receipt not found"})
}