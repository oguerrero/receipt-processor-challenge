package test

import (
	"receipt-processor-challenge/controllers"
	"receipt-processor-challenge/models"

	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.POST("/receipts/process", controllers.PostReceipt)
	router.GET("/receipts/:id/points", controllers.GetPoints)

	return router
}

func TestTargetReceipt(t *testing.T) {
	router := setupRouter()

	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
		},
		Total: "35.35",
	}
	jsonData, _ := json.Marshal(receipt)

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	id, ok := response["id"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, id)

	req, _ = http.NewRequest(http.MethodGet, "/receipts/"+id+"/points", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var pointsResponse map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	assert.NoError(t, err)

	points, ok := pointsResponse["points"].(float64)
	assert.True(t, ok)
	assert.Equal(t, float64(28), points)
}

func TestCornerMarketREceipt(t *testing.T) {
	router := setupRouter()

	receipt := models.Receipt{
		Retailer:     "M&M Corner Market",
		PurchaseDate: "2022-03-20",
		PurchaseTime: "14:33",
		Items: []models.Item{
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
			{ShortDescription: "Gatorade", Price: "2.25"},
		},
		Total: "9.00",
	}
	jsonData, _ := json.Marshal(receipt)

	req, _ := http.NewRequest(http.MethodPost, "/receipts/process", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	id, ok := response["id"].(string)
	assert.True(t, ok)
	assert.NotEmpty(t, id)

	req, _ = http.NewRequest(http.MethodGet, "/receipts/"+id+"/points", nil)
	w = httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var pointsResponse map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &pointsResponse)
	assert.NoError(t, err)

	points, ok := pointsResponse["points"].(float64)
	assert.True(t, ok)
	assert.Equal(t, float64(109), points)
}
