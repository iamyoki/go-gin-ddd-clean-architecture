package infrastructure

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	apperror "todo_api/app/error"
	"todo_api/modules/todo/domain"

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

type sqliteTodoRepository struct {
	db *gorm.DB
}

// FindById implements [domain.TodoRepositoryInterface].
func (s *sqliteTodoRepository) FindById(ctx context.Context, id uuid.UUID) (*domain.Todo, error) {
	todo, err := gorm.G[TodoEntity](s.db).Where("id = ?", id).First(ctx)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &apperror.NotFound{Msg: "todo id not found"}
		}

		return nil, err
	}

	return &domain.Todo{
		ID:          todo.ID,
		Title:       todo.Title,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
		CompletedAt: todo.CompletedAt,
	}, nil
}

// FindAll implements [domain.TodoRepositoryInterface].
func (s *sqliteTodoRepository) FindAll(ctx context.Context) ([]domain.Todo, error) {
	todoEntities, err := gorm.G[TodoEntity](s.db).Find(ctx)

	if err != nil {
		return nil, err
	}

	todos := make([]domain.Todo, 0, len(todoEntities))

	for _, e := range todoEntities {
		todos = append(todos, domain.Todo{
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

// Save implements [domain.TodoRepositoryInterface].
func (s *sqliteTodoRepository) Save(ctx context.Context, todo *domain.Todo) error {
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

func NewSqliteTodoRepository(db *gorm.DB) domain.TodoRepositoryInterface {
	return &sqliteTodoRepository{
		db: db,
	}
}
