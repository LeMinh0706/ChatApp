package repo

import (
	"context"

	"github.com/LeMinh0706/ChatApp/db"
)

type MessageRepo struct {
	queries *db.Queries
}

func NewMessageRepo(queries *db.Queries) (*MessageRepo, error) {
	return &MessageRepo{
		queries: queries,
	}, nil
}

func (m *MessageRepo) CreateMessage(ctx context.Context, arg db.CreateMessageParams) (db.Message, error) {
	return m.queries.CreateMessage(ctx, arg)
}
