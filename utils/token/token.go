package token

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type authCustomClaims struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(email string, id int64, role string) (string, error) {
	claims := &authCustomClaims{
		Id:    id,
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("123456789"))
	if err != nil {
		fmt.Println("Error signing token:", err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*authCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &authCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("123456789"), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	if claims, ok := token.Claims.(*authCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims or token is not valid")
}