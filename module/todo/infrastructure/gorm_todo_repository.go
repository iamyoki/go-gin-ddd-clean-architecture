package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"todo_api/app/apperror"
	"todo_api/module/todo/domain/todo"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TodoEntity struct {
	ID          uuid.UUID
	Title       string `gorm:"uniqueIndex;not null"`
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CompletedAt *time.Time
}

type gormTodoRepository struct {
	db *gorm.DB
}

// DeleteById implements [todo.TodoRepositoryInterface].
func (s *gormTodoRepository) DeleteById(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	var todo *todo.Todo

	err := s.db.Transaction(func(tx *gorm.DB) error {
		txRepo := *s
		txRepo.db = tx

		if t, err := txRepo.FindById(ctx, id); err != nil {
			return err
		} else {
			todo = t
		}

		if _, err := gorm.G[TodoEntity](tx).Where("id = ?", id).Delete(ctx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return todo, nil
}

// FindById implements [todo.TodoRepositoryInterface].
func (s *gormTodoRepository) FindById(ctx context.Context, id uuid.UUID) (*todo.Todo, error) {
	todoEntity, err := gorm.G[TodoEntity](s.db).Where("id = ?", id).First(ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &apperror.NotFound{Msg: "Todo id not found"}
		}

		return nil, err
	}

	return &todo.Todo{
		ID:          todoEntity.ID,
		Title:       todoEntity.Title,
		Completed:   todoEntity.Completed,
		CreatedAt:   todoEntity.CreatedAt,
		UpdatedAt:   todoEntity.UpdatedAt,
		CompletedAt: todoEntity.CompletedAt,
	}, nil
}

// FindAll implements [todo.TodoRepositoryInterface].
func (s *gormTodoRepository) FindAll(ctx context.Context) ([]todo.Todo, error) {
	todoEntities, err := gorm.G[TodoEntity](s.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	todos := make([]todo.Todo, 0, len(todoEntities))

	for _, e := range todoEntities {
		todos = append(todos, todo.Todo{
			ID:          e.ID,
			Title:       e.Title,
			Completed:   e.Completed,
			CreatedAt:   e.CreatedAt,
			UpdatedAt:   e.UpdatedAt,
			CompletedAt: e.CompletedAt,
		})
	}

	return todos, nil
}

// Save implements [todo.TodoRepositoryInterface].
func (s *gormTodoRepository) Save(ctx context.Context, todo *todo.Todo) error {
	todoEntity := TodoEntity{
		ID:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	}

	err := gorm.G[TodoEntity](s.db, clause.OnConflict{UpdateAll: true}).Create(ctx, &todoEntity)

	if err != nil && strings.Contains(err.Error(), "UNIQUE") {
		return &apperror.Conflict{Msg: fmt.Sprintf("The title `%s` already exists", todo.Title)}
	}

	return err
}

func NewGormTodoRepository(db *gorm.DB) todo.Repository {
	return &gormTodoRepository{
		db: db,
	}
}
