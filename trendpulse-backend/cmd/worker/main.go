package main

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/markiskorova/trendpulse-backend/internal/db"
	"github.com/markiskorova/trendpulse-backend/internal/tasks"
)

func main() {
	// Init DB
	db.InitDB()
	dbConn := db.Get()

	// Setup worker server
	server := asynq.NewServer(
		asynq.RedisClientOpt{Addr: "redis:6379"},
		asynq.Config{
			Concurrency: 10,
			Queues:      map[string]int{"default": 1},
		},
	)

	// Register task handler
	mux := asynq.NewServeMux()
	mux.HandleFunc("scrape:article", tasks.HandleScrapeArticleTask(dbConn))

	log.Println("⚙️ Worker started for 'scrape:article'")

	// Run server
	if err := server.Run(mux); err != nil {
		log.Fatalf("worker failed: %v", err)
	}
}
