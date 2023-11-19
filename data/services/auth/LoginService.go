package auth

import (
	"errors"
	"os"
)

type LoginService struct{}

func (s LoginService) Execute(login, password string) error {
	authLogin := os.Getenv("USER_LOGIN")
	authPassword := os.Getenv("USER_PASSWORD")

	if login != authLogin || password != authPassword {
		return errors.New("Credenciais inválidas")
	}

	return nil
}
