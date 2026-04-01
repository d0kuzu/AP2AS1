package order

import (
	"assign1/internal/config"
	"assign1/internal/database/models"
	"assign1/internal/database/repositories"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	cfg  *config.Settings
	repo *repositories.OrderRepository
}

func NewOrderHandler(cfg *config.Settings, repo *repositories.OrderRepository) *OrderHandler {
	return &OrderHandler{cfg: cfg, repo: repo}
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order := &models.Order{
		CustomerID: req.CustomerID,
		ItemName:   req.ItemName,
		Amount:     req.Amount,
	}

	createdOrder, err := h.repo.Create(c.Request.Context(), order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	paymentReq := PaymentRequest{
		OrderID: createdOrder.ID,
		Amount:  createdOrder.Amount,
	}

	paymentJSON, _ := json.Marshal(paymentReq)
	paymentResp, err := client.Post(fmt.Sprintf("%s/payments/", h.cfg.PaymentServiceURL), "application/json", bytes.NewBuffer(paymentJSON))

	newStatus := models.StatusFailed
	if err == nil && paymentResp.StatusCode == http.StatusOK {
		var pResp PaymentResponse
		if err := json.NewDecoder(paymentResp.Body).Decode(&pResp); err == nil {
			if pResp.Status == "Authorized" {
				newStatus = models.StatusPaid
			}
		}
		paymentResp.Body.Close()
	}

	_ = h.repo.UpdateStatus(c.Request.Context(), createdOrder.ID, newStatus)
	createdOrder.Status = newStatus

	c.JSON(http.StatusCreated, createdOrder)
}

func (h *OrderHandler) GetOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *OrderHandler) CancelOrder(c *gin.Context) {
	id := c.Param("id")
	order, err := h.repo.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}

	if order.Status != models.StatusPending {
		c.JSON(http.StatusBadRequest, gin.H{"error": "only pending orders can be cancelled"})
		return
	}

	err = h.repo.UpdateStatus(c.Request.Context(), id, models.StatusCancelled)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	order.Status = models.StatusCancelled
	c.JSON(http.StatusOK, order)
}
