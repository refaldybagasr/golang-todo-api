package repository

import (
	"context"
	"database/sql"
	"refaldybagasr/golang-todo-api/model/domain"
)

type ItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item
	Update(ctx context.Context, tx *sql.Tx, item domain.Item) domain.Item
	Delete(ctx context.Context, tx *sql.Tx, item domain.Item)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Item
}
