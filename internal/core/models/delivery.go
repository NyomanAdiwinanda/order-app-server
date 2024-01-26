package models

type Delivery struct {
	ID                uint `gorm:"primarykey"`
	OrderItemID       uint `gorm:"column:order_item_id" json:"order_item_id"`
	DeliveredQuantity int  `gorm:"column:delivered_quantity" json:"delivered_quantity"`
}
