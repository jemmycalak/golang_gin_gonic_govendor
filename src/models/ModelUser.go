package models

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id           string `json:"iduser"`
	Firstname    string `json:"firstname" binding:"required"`
	Lastname     string `json:"lastname" binding:"required"`
	Email        string `json:"email" binding:"required"`
	Password     string `json:"password" binding:"required"`
	ImageProfile string `json:"imageprofile" binding:"required"`
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

type LoginStruct struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type JwtStruct struct {
	UserId int `json:"userid"`
	jwt.StandardClaims
}
