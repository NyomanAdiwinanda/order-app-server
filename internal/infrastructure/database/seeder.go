package database

import (
	"github.com/NyomanAdiwinanda/order-app-server/internal/infrastructure/utils"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	// Read CSV files
	customerCompanyData, err := utils.ReadCSVFile("csv/customer_companies.csv")
	if err != nil {
		panic("Failed to read customer_companies.csv")
	}

	customerData, err := utils.ReadCSVFile("csv/customers.csv")
	if err != nil {
		panic("Failed to read customer.csv")
	}

	orderData, err := utils.ReadCSVFile("csv/orders.csv")
	if err != nil {
		panic("Failed to read orders.csv")
	}

	orderItemData, err := utils.ReadCSVFile("csv/order_items.csv")
	if err != nil {
		panic("Failed to read order_items.csv")
	}

	deliveryData, err := utils.ReadCSVFile("csv/deliveries.csv")
	if err != nil {
		panic("Failed to read deliveries.csv")
	}

	// Convert CSV data to struct slices
	customerCompanies := utils.CSVToCustomerCompanies(customerCompanyData)

	customers, err := utils.CSVToCustomers(customerData)
	if err != nil {
		panic("Failed to convert customers CSV")
	}

	orders := utils.CSVToOrders(orderData)

	orderItems, err := utils.CSVToOrderItems(orderItemData)
	if err != nil {
		panic(err)
	}

	deliveryItems, err := utils.CSVToDeliveries(deliveryData)
	if err != nil {
		panic(err)
	}

	// Seed data into the database
	for _, customerCompany := range customerCompanies {
		db.Create(&customerCompany)
	}
	for _, customer := range customers {
		db.Create(&customer)
	}
	for _, order := range orders {
		db.Create(&order)
	}
	for _, orderItem := range orderItems {
		db.Create(&orderItem)
	}
	for _, delivery := range deliveryItems {
		db.Create(&delivery)
	}
}
