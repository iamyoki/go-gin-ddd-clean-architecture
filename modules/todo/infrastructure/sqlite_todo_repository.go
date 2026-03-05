package infrastructure

import (
	"context"
	"strings"
	"time"
	apperror "todo_api/app/error"
	"todo_api/modules/todo/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoEntity struct {
	ID        uuid.UUID
	Title     string `gorm:"uniqueIndex;not null"`
	Completed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type sqliteTodoRepository struct {
	db *gorm.DB
}

// Save implements [domain.TodoRepositoryInterface].
func (s *sqliteTodoRepository) Save(ctx context.Context, todo *domain.Todo) error {
	todoEntity := TodoEntity{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	err := gorm.G[TodoEntity](s.db, clause.OnConflict{UpdateAll: true}).Create(ctx, &todoEntity)

	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		return &apperror.Conflict{Msg: "The `title` already exists"}
	}

	return err
}

func NewSqliteTodoRepository(db *gorm.DB) domain.TodoRepositoryInterface {
	return &sqliteTodoRepository{
		db: db,
	}
}
