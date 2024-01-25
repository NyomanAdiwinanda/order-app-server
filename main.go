package main

import (
	"github.com/NyomanAdiwinanda/order-app-server/database"
	"github.com/NyomanAdiwinanda/order-app-server/handlers"
	"github.com/NyomanAdiwinanda/order-app-server/repositories"
	"github.com/NyomanAdiwinanda/order-app-server/usecases"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	isSeeded := false
	db := database.InitializeDB()

	orderRepo := repositories.NewOrderRepository(db)
	orderUseCase := usecases.NewOrderUseCase(orderRepo)
	orderHandler := handlers.NewOrderHandler(orderUseCase)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.GET("/orders", orderHandler.GetAllOrders)
	r.GET("/migrate", func(c *gin.Context) {
		database.MigrateDB(db)
		c.JSON(200, gin.H{
			"message": "Database migration completed successfully",
		})
	})
	r.GET("/seed", func(c *gin.Context) {
		if isSeeded {
			c.JSON(400, gin.H{
				"message": "Database is already seeded",
			})
		} else {
			isSeeded = true
			database.SeedDB(db)
			c.JSON(200, gin.H{
				"message": "Database seeding completed successfully",
			})

		}
	})

	r.Run()
}
