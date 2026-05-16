package auth

import (
	"errors"
	"linksnap/cache"
	"linksnap/service"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

func Login(email string, password string, cache *cache.Cache) (string, error) {
	person, err := service.GetUserByEmail(cache, email)
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
	_, err := mail.ParseAddress(email)
	if err != nil {
		return errors.New("Invalid email address")
	}
	err = service.CreateUser(email, password, role)
	if err != nil {
		return err
	}
	return nil
}
