package utils

import (
	"time"

	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
)

func CSVToOrders(orderData [][]string) []models.Order {
	var Orders []models.Order
	const layout = "2006-01-02T15:04:05Z"
	location, _ := time.LoadLocation("Australia/Melbourne")

	for _, row := range orderData[1:] {
		var customerIdData string
		if row[3] == "ivan" {
			customerIdData = "1"
		} else {
			customerIdData = "2"
		}

		parsedTime, err := time.Parse(layout, row[1])
		if err != nil {
			panic("Error handling parse time")
		}
		melbourneTime := parsedTime.In(location)

		Order := models.Order{
			CreatedAt:  melbourneTime,
			OrderName:  row[2],
			CustomerID: customerIdData,
		}
		Orders = append(Orders, Order)
	}
	return Orders
}
