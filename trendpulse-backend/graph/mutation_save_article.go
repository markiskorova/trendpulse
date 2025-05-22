package graph

import (
	"context"
	"fmt"

	"trendpulse-backend/graph/model"
	"trendpulse-backend/internal/db/models"
	"trendpulse-backend/internal/queue"
)

func (r *mutationResolver) SaveArticle(ctx context.Context, url string) (*model.Article, error) {
	userID := getUserIDFromContext(ctx)
	if userID == 0 {
		return nil, fmt.Errorf("unauthorized")
	}

	// Save to DB
	article := models.Article{
		URL:    url,
		UserID: userID,
		Status: "pending",
	}
	if err := r.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	// Enqueue background job
	if err := queue.EnqueueScrapeTask(article.ID); err != nil {
		return nil, fmt.Errorf("failed to enqueue scrape task: %w", err)
	}

	// Return GraphQL model (optional transformation if needed)
	return &model.Article{
		ID:        fmt.Sprint(article.ID),
		URL:       article.URL,
		Status:    article.Status,
		CreatedAt: article.CreatedAt.String(),
	}, nil
}
