package utils

import (
	"strconv"

	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
)

func CSVToDeliveries(data [][]string) ([]models.Delivery, error) {
	var Deliveries []models.Delivery

	for _, row := range data[1:] {
		orderItemId, err := strconv.ParseUint(row[1], 10, 64)
		if err != nil {
			return nil, err
		}
		deliveredQuantity, err := strconv.Atoi(row[2])
		if err != nil {
			return nil, err
		}

		Delivery := models.Delivery{
			OrderItemID:       uint(orderItemId),
			DeliveredQuantity: int(deliveredQuantity),
		}
		Deliveries = append(Deliveries, Delivery)
	}
	return Deliveries, nil
}
