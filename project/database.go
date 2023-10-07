package main

import "fmt"

type Database struct {
	users    []User
	products []Product
}

func NewDatabase() Database {
	users := []User{
		{Email: "alexys.lozada@gmail.com", Password: "123456", IsAdmin: true},
		{Email: "pedro.perez@gmail.com", Password: "123456", IsAdmin: false},
	}
	products := []Product{
		{Name: "Arroz", Price: 12.24},
		{Name: "Papa", Price: 2.00},
		{Name: "Carne", Price: 25.40},
	}

	return Database{
		users:    users,
		products: products,
	}
}

func (d *Database) UserByEmailAndPassword(email, password string) (User, error) {
	for _, v := range d.users {
		if v.Email == email && v.Password == password {
			return v, nil
		}
	}

	return User{}, fmt.Errorf("email or password not valid")
}

func (d *Database) Products() []Product {
	return d.products
}

func (d *Database) ProductAdd(product Product) {
	d.products = append(d.products, product)
}
