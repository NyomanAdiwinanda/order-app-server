package utils

import "github.com/NyomanAdiwinanda/order-app-server/internal/core/models"

func CSVToCustomerCompanies(data [][]string) []models.CustomerCompany {
	var CustomerCompanies []models.CustomerCompany

	for _, row := range data[1:] {
		CustomerCompany := models.CustomerCompany{
			CompanyName: row[1],
		}
		CustomerCompanies = append(CustomerCompanies, CustomerCompany)
	}
	return CustomerCompanies
}
