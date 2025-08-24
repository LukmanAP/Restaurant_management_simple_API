package main

import (
	"fmt"
	"os"
	"restaurant-management/database"
	"restaurant-management/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Connect ke database
	if err := database.Connect(); err != nil {
		panic(fmt.Sprintf("Gagal koneksi database: %v", err))
	}
	defer database.DB.Close()

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	// router.Use(middleware.Aunthentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
