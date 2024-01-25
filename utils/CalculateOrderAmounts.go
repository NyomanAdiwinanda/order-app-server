package utils

import (
	"math"

	"github.com/NyomanAdiwinanda/order-app-server/models"
)

func CalculateOrderAmounts(orders []models.Order) {
	for i, order := range orders {
		var totalAmount, deliveredAmount float64
		for _, item := range order.OrderItems {
			totalAmount += item.PricePerUnit * float64(item.Quantity)
			deliveredAmount += item.PricePerUnit * float64(item.Delivery.DeliveredQuantity)
		}
		orders[i].TotalAmount = math.Round(totalAmount*100) / 100
		orders[i].DeliveredAmount = math.Round(deliveredAmount*100) / 100
	}
}
