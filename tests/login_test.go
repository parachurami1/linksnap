package tests

import (
	"linksnap/auth"
	"linksnap/db"
	"linksnap/storage"
	"testing"
)

func TestLogin(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := storage.CreateUser("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	tokn, err := auth.Login("Test@gmail.com", "P4ssw0rd")
	if err != nil || tokn == "" {
		t.Errorf("Could not Log in")
	}
	print("\n\n")
}
