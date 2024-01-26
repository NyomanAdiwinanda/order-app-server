package models

import "time"

type Order struct {
	ID              uint `gorm:"primarykey"`
	CreatedAt       time.Time
	OrderName       string      `gorm:"type:varchar;column:order_name" json:"order_name"`
	CustomerID      string      `gorm:"type:varchar;column:customer_id" json:"customer_id"`
	Customer        Customer    `json:"customer" gorm:"foreignKey:CustomerID"`
	OrderItems      []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
	TotalAmount     float64     `json:"total_amount" gorm:"-"`
	DeliveredAmount float64     `json:"delivered_amount" gorm:"-"`
}
