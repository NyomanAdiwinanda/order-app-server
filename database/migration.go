package database

import (
	"github.com/NyomanAdiwinanda/order-app-server/models"
	"gorm.io/gorm"
)

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.CustomerCompany{},
		&models.Customer{},
		&models.Order{},
		&models.OrderItem{},
		&models.Delivery{},
	)

	if err != nil {
		panic("Failed to migrate database!")
	}
}
