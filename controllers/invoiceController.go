package controllers

import "github.com/gin-gonic/gin"

func GetInvoices() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get all invoices")
	}
}

func GetInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Get an invoice by id")
	}
}

func CreateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Invoice created")
	}
}

func UpdateInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Invoice updated")
	}
}

func DeleteInvoice() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, "Invoice deleted")
	}
}
