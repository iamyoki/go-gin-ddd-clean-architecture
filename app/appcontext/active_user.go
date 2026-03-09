package appcontext

import (
	"context"

	"github.com/google/uuid"
)

type contextKey string

const activeUserKey contextKey = "activeUser"

type ActiveUser struct {
	ID    uuid.UUID
	Email string
}

func (a ActiveUser) IntoContext(parent context.Context) context.Context {
	return context.WithValue(parent, activeUserKey, a)
}

func GetActiveUser(ctx context.Context) (ActiveUser, bool) {
	a, ok := ctx.Value(activeUserKey).(ActiveUser)
	return a, ok
}
