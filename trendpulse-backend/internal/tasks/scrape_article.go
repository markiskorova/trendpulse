package tasks

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"

	"trendpulse-backend/internal/models"
)

type ScrapePayload struct {
	ArticleID uint `json:"article_id"`
}

func NewScrapeArticleTask(articleID uint) (*asynq.Task, error) {
	payload, err := json.Marshal(ScrapePayload{ArticleID: articleID})
	if err != nil {
		return nil, err
	}

	return asynq.NewTask("scrape:article", payload), nil
}

func HandleScrapeArticleTask(db *gorm.DB) asynq.HandlerFunc {
	return func(ctx context.Context, t *asynq.Task) error {
		var p ScrapePayload
		if err := json.Unmarshal(t.Payload(), &p); err != nil {
			return fmt.Errorf("failed to parse payload: %w", err)
		}

		var article models.Article
		if err := db.First(&article, p.ArticleID).Error; err != nil {
			return fmt.Errorf("article not found: %w", err)
		}

		resp, err := http.Get(article.URL)
		if err != nil {
			return fmt.Errorf("failed to fetch article URL: %w", err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			return fmt.Errorf("failed to parse HTML: %w", err)
		}

		article.Title = doc.Find("title").Text()
		article.Content = doc.Find("p").Text() // Simple version: you can improve this

		return db.Save(&article).Error
	}
}
