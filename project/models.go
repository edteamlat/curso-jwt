package main

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

type Product struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
