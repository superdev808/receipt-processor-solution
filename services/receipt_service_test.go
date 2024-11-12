package services

import (
	"main/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Test Case 1 - Basic Receipt",
			receipt: models.Receipt{
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
			},
			expected: 28, // Expected points as per example
		},
		{
			name: "Test Case 2 - Round Dollar Total",
			receipt: models.Receipt{
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
			},
			expected: 109, // Expected points as per example
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			points := CalculatePoints(test.receipt)
			if points != test.expected {
				t.Errorf("Expected %d points, but got %d", test.expected, points)
			}
		})
	}
}

func TestAlphanumericCount(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"Target", 6},
		{"M&M Corner Market", 14}, // '&' is not alphanumeric
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			count := alphanumericCount(test.input)
			if count != test.expected {
				t.Errorf("Expected %d alphanumeric characters, but got %d", test.expected, count)
			}
		})
	}
}

func TestIsRoundDollar(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"35.00", true},
		{"35.50", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isRoundDollar(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestIsMultipleOfQuarter(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"9.00", true},
		{"9.10", false},
	}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := isMultipleOfQuarter(test.input)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}

func TestDescriptionPoints(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
	}

	expected := 6 // 3 points for each item
	result := descriptionPoints(items)
	if result != expected {
		t.Errorf("Expected %d points, but got %d", expected, result)
	}
}
