package storage

import (
	"linksnap/db"
	"linksnap/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func SaveSlug(url string, slug string, user_id int) error {
	_, err := db.DB.Exec("INSERT INTO links(original_url,short_code,user_id,created_at,updated_at) VALUES($1,$2,$3,$4,$5)", url, slug, user_id, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(email string, password string, role string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.DB.Exec("INSERT INTO users(email,password_hash,created_at,role) VALUES($1,$2,$3,$4)", email, pwd, time.Now(), role)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(email string) (models.User, error) {
	var Person models.User
	err := db.DB.Get(&Person, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return models.User{}, err
	}
	return Person, nil
}

func CreateLink(link string, slug string, userID int) error {
	_, err := db.DB.Exec("INSERT INTO links(original_url,short_code,user_id,created_at,updated_at) VALUES($1,$2,$3,$4,$5)", link, slug, userID, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func GetLinks(userID int) ([]models.URL, error) {
	var links []models.URL
	err := db.DB.Select(&links, "SELECT * FROM links WHERE user_id='$1'", userID)
	if err != nil {
		return nil, err
	}
	return links, nil
}
