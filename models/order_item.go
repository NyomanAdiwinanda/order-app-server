package models

type OrderItem struct {
	ID           uint     `gorm:"primarykey"`
	OrderID      uint     `gorm:"column:order_id" json:"order_id"`
	PricePerUnit float64  `gorm:"type:decimal;column:price_per_unit" json:"price_per_unit"`
	Quantity     int      `gorm:"column:quantity" json:"quantity"`
	Product      string   `gorm:"type:varchar;column:product" json:"product"`
	Delivery     Delivery `json:"delivery" gorm:"foreignKey:OrderItemID"`
}
