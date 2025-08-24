package controllers

import "github.com/gin-gonic/gin"

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get an order by id")
	}
}

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all orders")
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order created")
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order updated")
	}
}

func DeleteOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Order deleted")
	}
}
