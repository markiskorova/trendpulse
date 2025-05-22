package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
	Articles []Article
}

type Article struct {
	gorm.Model
	URL     string
	Title   string
	Content string
	OwnerID uint // <- make sure this is uint
}
