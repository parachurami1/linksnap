package tests

import (
	"linksnap/cache"
	"linksnap/db"
	"linksnap/service"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := service.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	print("\n\n")
}

func TestGetUser(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := service.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	cac, err := cache.NewCache()
	if err != nil {
		t.Errorf("Failed to create cache")
	}
	_, err = service.GetUserByEmail(cac, "Test@gmail.com")
	if err != nil {
		t.Errorf("Could not get user")
	}
	print("\n\n")
}
