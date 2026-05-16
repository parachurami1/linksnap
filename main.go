package main

import "linksnap/db"

func main() {
	db.ConnectDB()
	db.RunMigrations("postgresql://postgres:p4ssw0rd@localhost:5432/Linksnap?sslmode=disable")

}
