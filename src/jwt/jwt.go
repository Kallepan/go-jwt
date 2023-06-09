package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gitlab.com/kallepan/go-jwt/env"
)

var jwtKey = []byte(env.GetValueFromEnv("JWT_KEY", "supersecret"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	UserId   string `json:"user_id"`
	jwt.RegisteredClaims
}

func GetUserIdFromToken(signedToken string) (userId string, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("failed to parse claims")
		return
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = errors.New("token expired")
		return
	}

	return claims.UserId, nil
}

func GenerateJWTToken(username string, email string, userId string) (tokenString string, err error) {
	expirationTime := time.Now().Add(12 * time.Hour)

	claims := &JWTClaim{
		Username: username,
		Email:    email,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{expirationTime},
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func ValidateJWTToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(*JWTClaim)

	if !ok {
		err = errors.New("failed to parse claims")
		return err
	}

	if claims.ExpiresAt.Time.Before(time.Now()) {
		err = errors.New("token expired")
		return err
	}

	return nil
}
