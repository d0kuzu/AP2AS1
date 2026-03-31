package api

import (
	"assign1/internal/api/payment"
	appModule "assign1/internal/app"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RouterStart(app *appModule.App) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Type", "Authorization"},
		MaxAge:       12 * 60 * 60,
	}))

	payment.PaymentRoutes(r, app)

	err := r.Run(":" + app.Cfg.HttpPort)
	if err != nil {
		log.Fatal("Router start error", err)
	}
}
