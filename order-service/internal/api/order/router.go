package order

import (
	appModule "assign1/internal/app"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(router *gin.Engine, app *appModule.App) {
	h := NewOrderHandler(app.Cfg, app.Repos.Order)

	orderGroup := router.Group("orders")
	{
		orderGroup.POST("/", h.CreateOrder)
		orderGroup.GET("/:id", h.GetOrder)
		orderGroup.PATCH("/:id/cancel", h.CancelOrder)
	}
}
