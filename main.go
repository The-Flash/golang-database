package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	City    string
	State   string
	Country string
	Pincode json.Number
}

type User struct {
	Name    string
	Age     json.Number
	Contact string
	Company string
	Address Address
}

func main() {
	dir := "./"
	db, err := New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John", "23", "123423", "Tech Company", Address{"Takoradi", "Maharashtra", "India", "400001"}},
		{"James", "23", "123423", "Tech Company", Address{"Accra", "Maharashtra", "India", "400001"}},
		{"Jenny", "23", "123423", "Tech Company", Address{"Kumasi", "Maharashtra", "India", "400001"}},
		{"Jill", "23", "123423", "Tech Company", Address{"Tema", "Maharashtra", "India", "400001"}},
	}

	for _, value := range employees {
		db.Write("users", value.Name, value)
	}

	records, err := db.ReadAll("users")
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(records)
}
