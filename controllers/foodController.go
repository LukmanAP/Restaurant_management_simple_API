package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all foods")
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get a food by id")
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Food created")
	}
}

// func round(num float64) int (

// )

// func toFixed(num float64, precision int) float64 {

// }

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Food updated")
	}
}

func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Food deleted")
	}
}
