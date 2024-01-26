package handlers

import (
	"math"
	"net/http"
	"strconv"

	"github.com/NyomanAdiwinanda/order-app-server/internal/core/models"
	"github.com/NyomanAdiwinanda/order-app-server/internal/core/usecases"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderUseCase usecases.OrderUseCase
}

func NewOrderHandler(orderUseCase usecases.OrderUseCase) *OrderHandler {
	return &OrderHandler{orderUseCase: orderUseCase}
}

func (h *OrderHandler) GetAllOrders(c *gin.Context) {
	page, pageSize, err := parsePaginationParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderName := c.Query("order_name")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	orders, totalCount, err := h.orderUseCase.GetAllOrders(page, pageSize, orderName, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error fetching orders",
			"data":    nil,
		})
		return
	}

	calculateOrderAmounts(orders)

	totalPage := (totalCount + pageSize - 1) / pageSize
	hasNextPage := page < totalPage

	c.JSON(http.StatusOK, gin.H{
		"status":        http.StatusOK,
		"message":       "Fetch orders success",
		"data":          orders,
		"has_next_page": hasNextPage,
		"total_page":    totalPage,
	})
}

func parsePaginationParams(c *gin.Context) (int, int, error) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		return 0, 0, err
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	if err != nil {
		return 0, 0, err
	}

	return page, pageSize, nil
}

func calculateOrderAmounts(orders []models.Order) {
	for i, order := range orders {
		var totalAmount, deliveredAmount float64
		for _, item := range order.OrderItems {
			totalAmount += item.PricePerUnit * float64(item.Quantity)
			deliveredAmount += item.PricePerUnit * float64(item.Delivery.DeliveredQuantity)
		}
		orders[i].TotalAmount = math.Round(totalAmount*100) / 100
		orders[i].DeliveredAmount = math.Round(deliveredAmount*100) / 100
	}
}
