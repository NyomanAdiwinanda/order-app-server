package utils

import (
	"strconv"

	"github.com/NyomanAdiwinanda/order-app-server/models"
)

func CSVToCustomers(data [][]string) ([]models.Customer, error) {
	var Customers []models.Customer

	for _, row := range data[1:] {
		companyID, err := strconv.ParseUint(row[4], 10, 64)
		if err != nil {
			return nil, err
		}

		Customer := models.Customer{
			IdName:      row[0],
			Login:       row[1],
			Password:    row[2],
			Name:        row[3],
			CompanyID:   uint(companyID),
			CreditCards: row[5],
		}
		Customers = append(Customers, Customer)
	}
	return Customers, nil
}
