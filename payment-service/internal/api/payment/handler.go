package payment

import (
	"assign1/internal/config"
	"assign1/internal/database/models"
	"assign1/internal/database/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PaymentHandler struct {
	cfg  *config.Settings
	repo *repositories.PaymentRepository
}

func NewPaymentHandler(cfg *config.Settings, repo *repositories.PaymentRepository) *PaymentHandler {
	return &PaymentHandler{cfg: cfg, repo: repo}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var req struct {
		OrderID string `json:"order_id" binding:"required"`
		Amount  int64  `json:"amount" binding:"required,gt=0"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	status := models.PaymentAuthorized
	if req.Amount > 100000 {
		status = models.PaymentDeclined
	}

	p := &models.Payment{
		OrderID:       req.OrderID,
		Amount:        req.Amount,
		Status:        status,
		TransactionID: uuid.New().String(),
	}

	p, err := h.repo.Create(c.Request.Context(), p)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (h *PaymentHandler) GetPaymentStatus(c *gin.Context) {
	orderID := c.Param("order_id")
	if orderID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "order_id is required"})
		return
	}

	p, err := h.repo.GetByOrderID(c.Request.Context(), orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}

	c.JSON(http.StatusOK, p)
}
