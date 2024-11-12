package services

import (
	"main/models"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func alphanumericCount(s string) int {
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	return len(re.FindAllString(s, -1))
}

func isRoundDollar(total string) bool {
	totalFloat, _ := strconv.ParseFloat(total, 64)
	return totalFloat == math.Floor(totalFloat)
}

func isMultipleOfQuarter(total string) bool {
	totalFloat, _ := strconv.ParseFloat(total, 64)
	return math.Mod(totalFloat, 0.25) == 0
}

func descriptionPoints(items []models.Item) int {
	points := 0
	for _, item := range items {
		descLen := len(strings.TrimSpace((item.ShortDescription)))
		if descLen%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

func isPurchaseDayOdd(dateStr string) bool {
	layout := "2006-01-02"
	date, _ := time.Parse(layout, dateStr)
	return date.Day()%2 == 1
}

func isAfternoonPurchase(timeStr string) bool {
	t, _ := time.Parse("15:04", timeStr)
	return t.Hour() >= 14 && t.Hour() < 16
}

func CalculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. Points for each alphanumeric character in the retailer name
	points += alphanumericCount(receipt.Retailer)

	// 2. 50 points if the total is a round dollar amount
	if isRoundDollar(receipt.Total) {
		points += 50
	}

	// 3. 25 points if the total is a multiple of 0.25
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// 4. 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Points for each item with description length multiple of 3
	points += descriptionPoints(receipt.Items)

	// 6. 6 points if purchase day is odd
	if isPurchaseDayOdd(receipt.PurchaseDate) {
		points += 6
	}

	// 7. 10 points if purchase time is between 2:00pm and 4:00pm
	if isAfternoonPurchase(receipt.PurchaseTime) {
		points += 10
	}

	return points
}
