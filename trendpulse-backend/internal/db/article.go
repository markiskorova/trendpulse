package db

import (
	"github.com/markiskorova/trendpulse-backend/internal/models"

	"gorm.io/gorm"
)

func SaveArticle(db *gorm.DB, article *models.Article) error {
	return db.Create(article).Error
}
