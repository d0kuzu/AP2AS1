package repositories

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type Chat struct {
	ID          string
	AssistantID string
	Customer    string
}

type Message struct {
	ID      string
	ChatID  string
	Content string
	Role    string
}

type ChatRepository struct {
	pool *pgxpool.Pool
}

func NewChatRepository(pool *pgxpool.Pool) *ChatRepository {
	return &ChatRepository{pool: pool}
}

func (r *ChatRepository) GetChatPage(assistantID string, page int32, limit int32) ([]Chat, error) {
	// TODO: implement with real query
	return []Chat{}, nil
}

func (r *ChatRepository) GetChatPagesCount(assistantID string, limit int32) (int32, error) {
	// TODO: implement with real query
	return 0, nil
}

func (r *ChatRepository) GetAllChatMessages(chatID string) ([]Message, error) {
	// TODO: implement with real query
	return []Message{}, nil
}

func (r *ChatRepository) SearchChatsByCustomer(assistantID string, searchTerm string) ([]Chat, int32, error) {
	// TODO: implement with real query
	return []Chat{}, 0, nil
}
