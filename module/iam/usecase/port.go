package usecase

type Hasher interface {
	HashPassowrd(password string) (string, error)
	VerifyPassword(hashedPassword string, password string) bool
}
