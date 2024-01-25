package usecases

import (
	"github.com/NyomanAdiwinanda/order-app-server/models"
	"github.com/NyomanAdiwinanda/order-app-server/repositories"
)

type OrderUseCase interface {
	GetAllOrders(page, pageSize int, orderName, product, startDate, endDate string) ([]models.Order, int, error)
}

type orderUseCase struct {
	orderRepo repositories.OrderRepository
}

func NewOrderUseCase(orderRepo repositories.OrderRepository) OrderUseCase {
	return &orderUseCase{orderRepo: orderRepo}
}

func (uc *orderUseCase) GetAllOrders(page, pageSize int, orderName, product, startDate, endDate string) ([]models.Order, int, error) {
	return uc.orderRepo.GetAllOrders(page, pageSize, orderName, product, startDate, endDate)
}
