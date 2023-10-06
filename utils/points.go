package utils

import (
	"receipt-processor-challenge/models"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// Count alphanumeric characters in Retailer name
	points += countAlphanumeric(receipt.Retailer)

	// Points for whole and quarter amounts
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 1.0) == 0 {
		points += 50
	}
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// Points for number of items
	points += (len(receipt.Items) / 2) * 5

	// Points for item descriptions with length divisible by 3
	for _, item := range receipt.Items {
		if len(strings.TrimSpace(item.ShortDescription)) % 3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Round((price * 0.2) + 0.5))
		}
	}

	// Points based on purchase date and time
	day, _ := strconv.Atoi(strings.Split(receipt.PurchaseDate, "-")[2])
	if day%2 != 0 {
		points += 6
	}

	time, _ := strconv.Atoi(strings.Split(receipt.PurchaseTime, ":")[0])
	if time > 13 && time < 16 {
		points += 10
	}

	return points
}

// Helper function to count alphanumeric characters in a string
func countAlphanumeric(s string) int {
	var count int
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			count++
		}
	}
	return count
}
