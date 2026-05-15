package tests

import (
	"linksnap/auth"
	"linksnap/db"
	"testing"
)

func TestRegister(t *testing.T) {
	db.ConnectDB()
	defer func() { db.DB.Exec("delete from users where email='Test@gmail.com'") }()
	err := auth.Register("Test@gmail.com", "P4ssw0rd", "user")
	if err != nil {
		t.Errorf("Could not create user")
	}
	print("\n\n")
}
