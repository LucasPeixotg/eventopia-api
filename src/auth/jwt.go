package auth

import (
	"fmt"
	"time"

	c "github.com/LucasPeixotg/eventopia-api/src/common"

	jwt "github.com/golang-jwt/jwt/v5"
)

type claims struct {
	c.Account
	jwt.RegisteredClaims
}

func signToken(acc *c.Account) (string, error) {

	// Create claims with multiple fields populated
	claims := claims{
		*acc,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwt_secret))
}

func verifyToken(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwt_secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return token, nil
}
