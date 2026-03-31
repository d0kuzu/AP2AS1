package payment

import (
	appModule "assign1/internal/app"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine, app *appModule.App) {
	h := NewPaymentHandler(app.Cfg, app.Repos.Payment)

	paymentGroup := router.Group("payments")
	{
		paymentGroup.POST("/process", h.ProcessPayment)
		paymentGroup.GET("/get_all", h.GetPayments)
	}
}
