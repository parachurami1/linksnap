package tests

import (
	"linksnap/cache"
	"linksnap/db"
	"linksnap/service"
	"testing"
)

func TestCreateLink(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	slug := service.Shorten("https://youtube.com")
	defer func() { db.DB.Exec("delete from urls where short_code=$1", slug) }()
	err := service.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	cac, err := cache.NewCache()
	if err != nil {
		t.Errorf("Failed to create cache")
	}
	err = service.CreateLink(cac, "https://youtube.com", slug, 1)
	if err != nil {
		t.Errorf("%s", err.Error())
	}
	print("\n\n")
}
