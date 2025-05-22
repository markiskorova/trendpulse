package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/markiskorova/trendpulse-backend/internal/models"
	"github.com/markiskorova/trendpulse-backend/internal/tasks"

	"github.com/gorilla/context"
	"github.com/hibiken/asynq"
	"gorm.io/gorm"
)

type SaveArticleRequest struct {
	URL string `json:"url"`
}

func SaveArticle(asynqClient *asynq.Client, dbConn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := context.Get(r, "user").(*models.User)

		var req SaveArticleRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		article := models.Article{
			URL:     req.URL,
			OwnerID: user.ID,
		}

		if err := dbConn.Create(&article).Error; err != nil {
			http.Error(w, "failed to save article", http.StatusInternalServerError)
			return
		}

		// ✅ Create scrape task
		task, err := tasks.NewScrapeArticleTask(article.ID)
		if err != nil {
			http.Error(w, "failed to build task", http.StatusInternalServerError)
			return
		}

		// ✅ Enqueue task
		if _, err := asynqClient.Enqueue(task); err != nil {
			http.Error(w, "failed to enqueue task", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(article)
	}
}

func GetArticles(dbConn *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := context.Get(r, "user").(*models.User)

		var articles []models.Article
		query := `SELECT * FROM articles WHERE owner_id = $1 ORDER BY created_at DESC`

		if err := dbConn.Select(&articles, query, user.ID); err != nil {
			http.Error(w, "failed to fetch articles", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(articles)
	}
}
