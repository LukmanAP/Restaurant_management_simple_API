package controllers

import "github.com/gin-gonic/gin"

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all tables")
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get a table by id")
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Table created")
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Table updated")
	}
}

func DeleteTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Table deleted")
	}
}
