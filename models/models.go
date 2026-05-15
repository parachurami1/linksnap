package models

import "time"

type URL struct {
	Id           int       `db:"id"`
	Url          string    `db:"original_url"`
	Short_code   string    `db:"short_code"`
	UserID       string    `db:"user_id"`
	Clicks       int       `db:"clicks"`
	Time_created time.Time `db:"created_at"`
	Last_updated time.Time `db:"updated_at"`
}

type User struct {
	Id         int       `db:"id"`
	Email      string    `db:"email"`
	Password   string    `db:"password_hash"`
	Role       string    `db:"role"`
	Created_at time.Time `db:"created_at"`
}
