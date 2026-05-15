package auth

import (
	"linksnap/storage"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string) (string, error) {
	person, err := storage.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(person.Password), []byte(password))
	if err != nil {
		return "", err
	}
	token, err := GenerateToken(person.Id, person.Role)
	if err != nil {
		return "", err
	}
	return token, nil
}

func Register(email string, password string, role string) error {
	err := storage.CreateUser(email, password, role)
	if err != nil {
		return err
	}
	return nil
}
