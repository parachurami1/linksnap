package storage

import (
	"linksnap/cache"
	"linksnap/db"
	"linksnap/models"
	"time"
)

func SaveSlug(url string, slug string, user_id int) error {
	_, err := db.DB.Exec("INSERT INTO links(original_url,short_code,user_id,created_at,updated_at) VALUES($1,$2,$3,$4,$5)", url, slug, user_id, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func SaveUser(email string, password string, role string) error {
	_, err := db.DB.Exec("INSERT INTO users(email,password_hash,created_at,role) VALUES($1,$2,$3,$4)", email, password, time.Now(), role)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo(email string) (models.User, error) {
	var Person models.User
	err := db.DB.Get(&Person, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		return models.User{}, err
	}
	return Person, nil
}

func SaveLink(link string, slug string, userID int) error {
	_, err := db.DB.Exec("INSERT INTO urls(original_url,short_code,user_id,created_at,updated_at) VALUES($1,$2,$3,$4,$5)", link, slug, userID, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func GetLinks(userID int) ([]models.URL, error) {
	var links []models.URL
	err := db.DB.Select(&links, "SELECT * FROM urls WHERE user_id=$1", userID)
	if err != nil {
		return nil, err
	}
	return links, nil
}
func GetLinksBySlug(slug string) (models.URL, error) {
	var url models.URL
	err := db.DB.Get(&url, "SELECT * FROM urls WHERE short_code=$1", slug)
	if err != nil {
		return models.URL{}, err
	}
	return url, nil
}

func UpdateCount(cache *cache.Cache, slug string) error {
	cache.DeleteByPattern("links:")
	_, err := db.DB.Exec("UPDATE urls SET clicks=clicks + 1 WHERE short_code=$1", slug)
	if err != nil {
		return err
	}
	return nil
}
