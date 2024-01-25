package utils

import (
	"strconv"

	"github.com/NyomanAdiwinanda/order-app-server/models"
)

func CSVToOrderItems(data [][]string) ([]models.OrderItem, error) {
	var OrderItems []models.OrderItem

	for _, row := range data[1:] {
		orderId, err := strconv.ParseUint(row[1], 10, 64)
		if err != nil {
			return nil, err
		}

		var pricePerUnit float64
		if row[2] != "" {
			pricePerUnit, err = strconv.ParseFloat(row[2], 64)
			if err != nil {
				pricePerUnit = 0
			}
		}

		quantity, err := strconv.Atoi(row[3])
		if err != nil {
			return nil, err
		}

		OrderItem := models.OrderItem{
			OrderID:      uint(orderId),
			PricePerUnit: float64(pricePerUnit),
			Quantity:     int(quantity),
			Product:      row[4],
		}
		OrderItems = append(OrderItems, OrderItem)
	}
	return OrderItems, nil
}
