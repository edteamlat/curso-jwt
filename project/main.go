package main

import "net/http"

func main() {
	secret := "ZXN0YWVzdW5hY2xhdmVtdXlzZWd1cmFwYXJhZ2VuZXJhcnRva2Vuc2p3dGNvbmdvZW5FRHRlYW0="
	db := NewDatabase()
	jwt := NewJWT(secret)

	handler := NewHandler(&db, jwt)
	middleware := NewMiddleware(jwt)

	http.HandleFunc("/login", handler.Login)
	http.HandleFunc("/products", middleware.ValidateJWT(handler.Products))
	http.HandleFunc("/create-product", middleware.ValidateJWTIsAdmin(handler.ProductCreate))

	http.ListenAndServe(":8080", nil)
}
