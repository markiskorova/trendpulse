package models

import "time"

type Article struct {
	ID        uint   `gorm:"primaryKey"`
	URL       string `gorm:"not null"`
	Status    string `gorm:"default:pending"`
	Content   string
	UserID    uint
	CreatedAt time.Time
}
