package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, userRole string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    userRole,
		"exp":     time.Now().Add(1 * time.Hour).Unix(),
	}
	tokn := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokn.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return token, err
}

func ParseToken(token string) (*jwt.Token, error) {
	// 1. Capture the return values from jwt.Parse
	result, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		// 2. Use a type assertion, not == comparison
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		// 3. []byte(...) not byte[...]  — it's a slice conversion, not an index
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
