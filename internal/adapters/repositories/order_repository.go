package repositories

import (
	"strings"

	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
	"gorm.io/gorm"
)

type OrderRepository interface {
	GetAllOrders(page, pageSize int, orderName, startDate, endDate string) ([]models.Order, int, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db: db}
}

func (r *orderRepository) GetAllOrders(page, pageSize int, orderName, startDate, endDate string) ([]models.Order, int, error) {
	var orders []models.Order
	var totalCount int64

	query := r.db.Model(&models.Order{})

	if orderName != "" {
		orderName = strings.ToLower(orderName)
		subQuery := r.db.Model(&models.OrderItem{}).Select("order_id").Where("LOWER(product) LIKE ?", "%"+orderName+"%").Group("order_id")
		query = query.Where("LOWER(order_name) LIKE ? OR id IN (?)", "%"+orderName+"%", subQuery)
	}

	if startDate != "" {
		query = query.Where("created_at >= ?", startDate)
	}
	if endDate != "" {
		query = query.Where("created_at <= ?", endDate)
	}

	if err := query.Count(&totalCount).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	if err := query.Preload("OrderItems").
		Preload("OrderItems.Delivery").
		Preload("Customer").
		Preload("Customer.Company").
		Offset(offset).
		Limit(pageSize).
		Find(&orders).Error; err != nil {
		return nil, 0, err
	}

	return orders, int(totalCount), nil
}
