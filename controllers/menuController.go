package controllers

import "github.com/gin-gonic/gin"

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all menus")
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get a menu by id")
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Menu created")
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Menu updated")
	}
}

func DeleteMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Menu deleted")
	}
}
