package main

import (
  "fmt"

  jwt "github.com/dgrijalva/jwt-go"
)

var (
  tokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjoiQWxleHlzIExvemFkYSIsImVtYWlsIjoiYWxleHlzLmxvemFkYUBnbWFpbC5jb20iLCJpYXQiOjE2OTY2OTY3MDAsImV4cCI6MTY5Njc4MzEwMH0.sSrVQvTHdNaOPp5cql27Ya_vssyckUEocH0D4NDqYgE"
  secret = []byte("this-is-a-very-long-and-secure-secret")
)

func validateMethodAndGetSecret(token *jwt.Token) (any, error) {
  _, ok := token.Method.(*jwt.SigningMethodHMAC)
  if !ok {
    return nil, fmt.Errorf("Método de la firma no es válido")
  }

  return secret, nil
}

func Validate() {
  token, err := jwt.Parse(tokenString, validateMethodAndGetSecret)
  if err != nil {
    fmt.Printf("token no valido, razón: %v", err)

    return
  }


  userData, ok := token.Claims.(jwt.MapClaims)
  if !ok {
    fmt.Println("no se pudo obtener la información del payload")
    return
  }

  fmt.Println("user", userData["user"])
  fmt.Println("email", userData["email"])
}

