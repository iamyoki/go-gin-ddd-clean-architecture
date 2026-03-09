package user

import (
	"net/mail"
	"regexp"
	"time"
	"todo_api/app/appcontext"
	"todo_api/app/apperror"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

var (
	pwReg                        = regexp.MustCompile(`^[a-zA-Z0-9]{6,20}$`)
	pwError *apperror.Validation = &apperror.Validation{Msg: "Invalid password format, only allow a-z, A-Z, 0-9 combination within 6-20 characters"}
)

func ValidatePassword(password string) error {
	if pwReg.MatchString(password) {
		return nil
	}
	return pwError
}

func Create(email string, hashedPassword string) (*User, error) {
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, &apperror.Validation{Msg: "Invalid email"}
	}

	now := time.Now()
	return &User{
		ID:             uuid.New(),
		Email:          email,
		HashedPassword: hashedPassword,
		CreatedAt:      now,
		UpdatedAt:      now,
	}, nil
}

func (user *User) ToActiveUser() appcontext.ActiveUser {
	return appcontext.ActiveUser{
		ID:    user.ID,
		Email: user.Email,
	}
}
