package usecases

import (
	"github.com/NyomanAdiwinanda/order-app-server/internal/adapters/repositories"
	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
)

type OrderUseCase interface {
	GetAllOrders(page, pageSize int, orderName, startDate, endDate string) ([]models.Order, int, error)
}

type orderUseCase struct {
	orderRepo repositories.OrderRepository
}

func NewOrderUseCase(orderRepo repositories.OrderRepository) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}

func (uc *orderUseCase) GetAllOrders(page, pageSize int, orderName, startDate, endDate string) ([]models.Order, int, error) {
	return uc.orderRepo.GetAllOrders(page, pageSize, orderName, startDate, endDate)
}
