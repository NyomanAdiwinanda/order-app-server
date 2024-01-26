package database

import (
	"github.com/NyomanAdiwinanda/order-app-server/internal/infrastructure/utils"
	"gorm.io/gorm"
)

func SeedDB(db *gorm.DB) {
	// Read CSV file
	customerCompanyData, err := utils.ReadCSVFile("csv/customer_companies.csv")
	if err != nil {
		panic("Failed to read customer_companies.csv")
	}

	// Convert CSV data to struct slices
	customerCompanies := utils.CSVToCustomerCompanies(customerCompanyData)

	// Seed data into the database
	for _, customerCompany := range customerCompanies {
		db.Create(&customerCompany)
	}

	// Repeat same process for other CSV files
	customerData, err := utils.ReadCSVFile("csv/customers.csv")
	if err != nil {
		panic("Failed to read customer.csv")
	}

	customers, err := utils.CSVToCustomers(customerData)
	if err != nil {
		panic("Failed to convert customers CSV")
	}

	for _, customer := range customers {
		db.Create(&customer)
	}

	orderData, err := utils.ReadCSVFile("csv/orders.csv")
	if err != nil {
		panic("Failed to read orders.csv")
	}

	orders, err := utils.CSVToOrders(db, orderData)
	if err != nil {
		panic("Failed to convert orders CSV")
	}

	for _, order := range orders {
		db.Create(&order)
	}

	orderItemData, err := utils.ReadCSVFile("csv/order_items.csv")
	if err != nil {
		panic("Failed to read order_items.csv")
	}

	orderItems, err := utils.CSVToOrderItems(orderItemData)
	if err != nil {
		panic(err)
	}

	for _, orderItem := range orderItems {
		db.Create(&orderItem)
	}

	deliveryData, err := utils.ReadCSVFile("csv/deliveries.csv")
	if err != nil {
		panic("Failed to read deliveries.csv")
	}

	deliveryItems, err := utils.CSVToDeliveries(deliveryData)
	if err != nil {
		panic(err)
	}

	for _, delivery := range deliveryItems {
		db.Create(&delivery)
	}
}
