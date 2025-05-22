package graph

import (
	"context"
	"strconv"

	"trendpulse-backend/graph/generated"
	"trendpulse-backend/graph/model"
	"trendpulse-backend/internal/models"
	"trendpulse-backend/internal/tasks"

	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type Resolver struct {
	DB    *gorm.DB
	Asynq *asynq.Client
}

// --- Root Interfaces ---
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// --- Query Resolver ---
type queryResolver struct{ *Resolver }

func (r *queryResolver) MyArticles(ctx context.Context) ([]*model.Article, error) {
	userID := ctx.Value("user_id").(uint)

	var articles []models.Article
	if err := r.DB.Where("owner_id = ?", userID).Find(&articles).Error; err != nil {
		return nil, err
	}

	var result []*model.Article
	for _, a := range articles {
		title := a.Title
		content := a.Content
		result = append(result, &model.Article{
			ID:      strconv.Itoa(int(a.ID)),
			URL:     a.URL,
			Title:   &title,
			Content: &content,
		})
	}

	return result, nil
}

// --- Mutation Resolver ---
type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SaveArticle(ctx context.Context, url string) (*model.Article, error) {
	userID := ctx.Value("user_id").(uint)

	article := models.Article{
		URL:     url,
		OwnerID: userID,
	}

	if err := r.DB.Create(&article).Error; err != nil {
		return nil, err
	}

	task, err := tasks.NewScrapeArticleTask(article.ID)
	if err != nil {
		return nil, err
	}

	if _, err := r.Asynq.Enqueue(task); err != nil {
		return nil, err
	}

	title := article.Title
	content := article.Content
	return &model.Article{
		ID:      strconv.Itoa(int(article.ID)),
		URL:     article.URL,
		Title:   &title,
		Content: &content,
	}, nil
}
