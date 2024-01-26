package utils

import (
	"strconv"
	"time"

	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
	"gorm.io/gorm"
)

func CSVToOrders(db *gorm.DB, orderData [][]string) ([]models.Order, error) {
	var Orders []models.Order
	const layout = "2006-01-02T15:04:05Z"
	location, _ := time.LoadLocation("Australia/Melbourne")

	for _, row := range orderData[1:] {
		var customer models.Customer
		if err := db.Where("id_name = ?", row[3]).First(&customer).Error; err != nil {
			return nil, err
		}

		parsedTime, err := time.Parse(layout, row[1])
		if err != nil {
			return nil, err
		}
		melbourneTime := parsedTime.In(location)

		Order := models.Order{
			CreatedAt:  melbourneTime,
			OrderName:  row[2],
			CustomerID: strconv.Itoa(int(customer.ID)),
		}
		Orders = append(Orders, Order)
	}
	return Orders, nil
}
