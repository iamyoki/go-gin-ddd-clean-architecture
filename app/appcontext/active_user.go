package appcontext

import "github.com/google/uuid"

type ActiveUser struct {
	ID    uuid.UUID
	Email string
}
