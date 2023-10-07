package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JWT struct {
	secret string
}

func NewJWT(secret string) JWT {
	return JWT{secret: secret}
}

func (j JWT) Create(user User) (string, error) {
	now := time.Now()
	iat := now.Unix()
	eat := now.AddDate(0, 0, 1).Unix()

	payload := jwt.MapClaims{
		"email":    user.Email,
		"is_admin": user.IsAdmin,
		"iat":      iat,
		"exp":      eat,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(j.secret))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}

func (j JWT) Validate(tokenString string) (User, error) {
	token, err := jwt.Parse(tokenString, j.validateMethodAndGetSecret)
	if err != nil {
		fmt.Printf("token no valido, razón: %v\n", err)

		return User{}, err
	}

	userData, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("no se pudo obtener la información del payload")
		return User{}, err
	}

	resp := User{
		Email:   userData["email"].(string),
		IsAdmin: userData["is_admin"].(bool),
	}

	return resp, nil
}

func (j JWT) validateMethodAndGetSecret(token *jwt.Token) (any, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, fmt.Errorf("Método de la firma no es válido")
	}

	return []byte(j.secret), nil
}
