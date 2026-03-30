package chat

import (
	"assign1/internal/config"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"assign1/internal/database/repositories"
)

type ChatHandler struct {
	cfg  *config.Settings
	repo *repositories.ChatRepository
}

func NewChatHandler(cfg *config.Settings, repo *repositories.ChatRepository) *ChatHandler {
	return &ChatHandler{cfg: cfg, repo: repo}
}

func (h *ChatHandler) GetAllChats(c *gin.Context) {
	pageStr := c.DefaultQuery("page", "1")
	page, err := strconv.ParseInt(pageStr, 10, 32)
	if err != nil || page < 1 {
		page = 1
	}

	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	chatsPerPage := int32(10)

	chats, err := h.repo.GetChatPage(assistantID, int32(page), chatsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer": chats,
	})
}

func (h *ChatHandler) GetPagination(c *gin.Context) {
	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	chatsPerPage := int32(10)
	pagesCount, err := h.repo.GetChatPagesCount(assistantID, chatsPerPage)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch pagination", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": pagesCount})
}

func (h *ChatHandler) GetChat(c *gin.Context) {
	chatID := c.Query("chat")
	if chatID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat parameter is required (chatID)"})
		return
	}

	messages, err := h.repo.GetAllChatMessages(chatID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch chat messages", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"answer": gin.H{
		"chat_id":  chatID,
		"messages": messages,
		"count":    len(messages),
	}})
}

func (h *ChatHandler) SearchChat(c *gin.Context) {
	searchTerm := c.Query("chat")
	assistantID := c.Query("assistant_id")
	if assistantID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "assistant_id is required"})
		return
	}

	if searchTerm == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "chat parameter is required (search term)"})
		return
	}

	chats, totalCount, err := h.repo.SearchChatsByCustomer(assistantID, searchTerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to search chats", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"answer":      chats,
		"total_count": totalCount,
	})
}
