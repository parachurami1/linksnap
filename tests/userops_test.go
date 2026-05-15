package tests

import (
	"linksnap/db"
	"linksnap/storage"
	"testing"
)

func TestCreateUser(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := storage.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	print("\n\n")
}

func TestGetUser(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := storage.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	_, err = storage.GetUserByEmail("Test@gmail.com")
	if err != nil {
		t.Errorf("Could not get user")
	}
	print("\n\n")
}
