package controllers

import "github.com/gin-gonic/gin"

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all order items")
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get an order item by id")
	}
}

// func GetOrderItemsByOrder(id string) ([]OrderItem, error) {

// }

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order item created")
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order item updated")
	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order item deleted")
	}
}
