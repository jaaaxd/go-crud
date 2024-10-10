package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title          string  `json:"title" binding:"required"`
	Subtitle       *string `json:"subtitle"`
	Desc           *string `json:"desc"`
	Price          uint    `json:"price" binding:"required"`
	GuruInfo       *string `json:"guru_info"`
	Type           string  `json:"type" binding:"required"`
	RelatedStock   *string `json:"related_stock"`
	ExpectedReturn *string `json:"expected_return"`
}

type User struct {
	gorm.Model
	Email       string    `json:"email" binding:"required" gorm:"unique"`
	Password    string    `json:"password" binding:"required"`
	Firstname   string    `json:"firstname" binding:"required"`
	Lastname    string    `json:"lastname" binding:"required"`
	Experience  string    `json:"experience" binding:"required"`
	Type        string    `json:"type" binding:"required"`
	PhoneNumber string    `json:"phone_number" binding:"required"`
	Birthday    time.Time `json:"birthday" binding:"required"`
}