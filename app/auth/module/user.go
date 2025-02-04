package module

import (
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type MyClaims struct {
	Id int `json:"id"`
	jwt.RegisteredClaims
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Salt     string `json:"salt"`
}
