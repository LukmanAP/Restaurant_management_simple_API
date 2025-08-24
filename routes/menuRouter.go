package routes

import (
	"restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(IncomingRoutes *gin.Engine) {
	IncomingRoutes.GET("/menus", controllers.GetMenus())
	IncomingRoutes.GET("/menus/:menu_id", controllers.GetMenu())
	IncomingRoutes.POST("/menus", controllers.CreateMenu())
	IncomingRoutes.PATCH("/menus/:menu_id", controllers.UpdateMenu())
	IncomingRoutes.DELETE("/menus/:menu_id", controllers.DeleteMenu())
}
