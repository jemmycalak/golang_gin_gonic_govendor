package models

import (
	"time"
)

type User struct {
	Id           string `json:"iduser"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	ImageProfile string `json:"imageprofile"`
	CreateAt     time.Time
	UpdateAt     time.Time
}

type Users []User

func NewUser() *User {
	return &User{
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
}
