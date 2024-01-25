package handlers

import (
	"net/http"
	"strconv"

	"github.com/NyomanAdiwinanda/order-app-server/usecases"
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
	product := c.Query("product")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	orders, totalCount, err := h.orderUseCase.GetAllOrders(page, pageSize, orderName, product, startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Error fetching orders",
			"data":    nil,
		})
		return
	}

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
