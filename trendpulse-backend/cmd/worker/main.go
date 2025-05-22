package main

import (
	"log"

	"trendpulse-backend/internal/db"
	"trendpulse-backend/internal/tasks"

	"github.com/hibiken/asynq"
)

func main() {
	db.InitDB()
	dbConn := db.Get()

	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "redis:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues:      map[string]int{"default": 1},
		},
	)

	mux := asynq.NewServeMux()
	mux.HandleFunc("scrape:article", tasks.HandleScrapeArticleTask(dbConn))

	log.Println("⚙️ Worker started for 'scrape:article'")
	if err := server.Run(mux); err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
