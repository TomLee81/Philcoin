package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Claims extends StandardClaims with custom fields
type Claims struct {
	jwt.StandardClaims
}

// GenerateToken creates a signed JWT for a given userID
func GenerateToken(userID string, secret string, expiry time.Duration) (string, error) {
	claims := &Claims{
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Add(expiry).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// ValidateToken parses and validates a JWT string
func ValidateToken(tokenStr string, secret string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
