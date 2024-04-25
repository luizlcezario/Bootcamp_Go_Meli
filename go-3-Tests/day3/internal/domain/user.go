package domain

import (
	"time"

	"github.com/luizlcezario/MelicampGoWeb/cmd/server/utils"
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

func NewUser(name string, sourName string, email string, age int, height float32, isActive bool) *User {
	return &User{
		Id:         utils.GeneratedId(email + name + time.Now().String()),
		Name:       name,
		SourName:   sourName,
		Email:      email,
		Age:        age,
		Heigth:     height,
		IsActive:   isActive,
		CreateadAt: time.Now(),
	}
}
