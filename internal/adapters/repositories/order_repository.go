package repositories

import (
	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(page, pageSize int, orderName, product, startDate, endDate string) ([]models.Order, int, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAllOrders(page, pageSize int, orderName, product, startDate, endDate string) ([]models.Order, int, error) {
	var orders []models.Order
	var totalCount int64
	offset := (page - 1) * pageSize
	query := r.db

	// Implement filtering logic
	if orderName != "" {
		query = query.Where("order_name LIKE ?", "%"+orderName+"%")
	}
	if product != "" {
		query = query.Where("product LIKE ?", "%"+product+"%") // Adjust based on your database schema
	}
	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Model(&models.Order{}).Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Preload("OrderItems.Delivery").Preload("Customer").Preload("Customer.Company").Offset(offset).Limit(pageSize).Find(&orders).Error; err != nil {
		return nil, 0, err
	}
	return orders, int(totalCount), nil
}
