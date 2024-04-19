package main

import (
	"errors"
	"time"
)

type User struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	SourName   string    `json:"sour_name"`
	Email      string    `json:"email"`
	Age        int       `json:"age"`
	Heigth     float32   `json:"height"`
	IsActive   bool      `json:"is_active"`
	CreateadAt time.Time `json:"created_at"`
}

var Users []*User = []*User{
	NewUser("luiz", "lima", "luiz@gmail", 20, 1.76, true),
	NewUser("milena", "carecho", "care@gmail", 22, 1.56, true),
	NewUser("luan", "cardoso", "lcar@gmail", 23, 1.67, false),
	NewUser("felipe", "lima", "fel@gmail", 10, 1.90, true),
}

func NewUser(name string, sourName string, email string, age int, height float32, isActive bool) *User {
	return &User{
		Id:         generatedId(email + name + time.Now().String()),
		Name:       name,
		SourName:   sourName,
		Email:      email,
		Age:        age,
		Heigth:     height,
		IsActive:   isActive,
		CreateadAt: time.Now(),
	}
}

func (user *User) GetQueryParametersValids() map[string]int {
	return map[string]int{"name": 2, "sourname": 3, "email": 4, "age": 5, "height": 6, "isactive": 7}
}

func findById(users []*User, id string) (User, error) {
	for _, e := range users {
		if e.Id == id {
			return *e, nil
		}
	}
	return User{}, errors.New("Nao foi possivel achar o usuario com o Id: " + id)
}
