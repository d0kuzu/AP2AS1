package order

import (
	appModule "assign1/internal/app"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, app *appModule.App) {
	h := NewOrderHandler(app.Cfg, app.Repos.Order)

	orderGroup := router.Group("orders")
	{
		orderGroup.POST("/create", h.CreateOrder)
		orderGroup.GET("/get_all", h.GetOrders)
	}
}
