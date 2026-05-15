package service

import "math/rand/v2"

func Shorten(url string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	slug := ""
	for range 5 {
		slug += string(chars[rand.IntN(len(chars))])
	}
	return slug
}

func S() {

}
