package service

import (
	"context"
	"encoding/json"
	"fmt"
	"linksnap/cache"
	"linksnap/models"
	"linksnap/storage"
	"log/slog"
	"math/rand/v2"
	"time"

	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

func Shorten(url string) string {
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	slug := ""
	for range 5 {
		slug += string(chars[rand.IntN(len(chars))])
	}
	return slug
}

func CreateUser(email string, password string, role string) error {
	pwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	err = storage.SaveUser(email, string(pwd), role)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByEmail(cache *cache.Cache, email string) (models.User, error) {
	var Person models.User
	val := cache.Client.Get(context.Background(), fmt.Sprintf("user:%v", email))
	res, err := val.Result()
	if err == redis.Nil || err != nil {
		if err != nil && err != redis.Nil {
			slog.Error("Redis get failed", "error", err)
		}
		Person, err := storage.GetUserInfo(email)
		if err != nil {
			return models.User{}, err
		}
		go func(usr *models.User) {
			data, err := json.Marshal(usr)
			if err != nil {
				return
			}
			cache.Client.Set(context.Background(), fmt.Sprintf("user:%v", email), data, 15*time.Minute)
		}(&Person)

		return Person, nil
	}
	err = json.Unmarshal([]byte(res), &Person)
	return Person, nil
}

func GetLinksByID(cache *cache.Cache, userID int) ([]models.URL, error) {
	val := cache.Client.Get(context.Background(), fmt.Sprintf("links:%v", userID))
	res, err := val.Result()
	var links []models.URL

	if err == redis.Nil || err != nil {
		if err != nil && err != redis.Nil {
			slog.Error("Error getting from redis", "Error", err)
		}
		links, err := storage.GetLinks(userID)
		if err != nil {
			return nil, err
		}
		go func(lnk []models.URL) {
			data, err := json.Marshal(&lnk)
			if err != nil {
				return
			}
			cache.Client.Set(context.Background(), fmt.Sprintf("links:%v", userID), data, 15*time.Minute)
		}(links)
		return links, nil
	}

	err = json.Unmarshal([]byte(res), &links)
	if err != nil {
		return nil, err
	}
	return links, nil
}

func CreateLink(link string, slug string, userID int) error {
	err := storage.SaveLink(link, slug, userID)
	if err != nil {
		return err
	}
	return nil
}
