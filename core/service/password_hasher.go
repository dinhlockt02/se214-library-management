package coreservice

type PasswordHasher interface {
	HashPassword(password string) (string, error)
	VerifyPassword(rawPassword string, hashedPassword string) (bool, error)
}
