package auth

import (
	"errors"
	"linksnap/models"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userID int, userRole string) (string, error) {
	claims := models.Claim{
		User_ID: userID,
		Role:    userRole,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
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

func ParseTokenClaims(token string) (*models.Claim, error) {
	tokn, err := jwt.ParseWithClaims(token, &models.Claim{}, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tokn.Claims.(*models.Claim)
	if !ok || !tokn.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
