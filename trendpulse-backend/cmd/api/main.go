package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/hibiken/asynq"

	"github.com/markiskorova/trendpulse-backend/graph"
	"github.com/markiskorova/trendpulse-backend/graph/generated"
	"github.com/markiskorova/trendpulse-backend/internal/db"
	"github.com/markiskorova/trendpulse-backend/internal/handlers"
	"github.com/markiskorova/trendpulse-backend/internal/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	// Init DB
	db.InitDB()
	dbConn := db.Get()

	// Init Asynq client
	asynqClient := asynq.NewClient(asynq.RedisClientOpt{Addr: "redis:6379"})

	// Router
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
	r.HandleFunc("/register", handlers.Register(dbConn)).Methods("POST")
	r.HandleFunc("/login", handlers.Login(dbConn)).Methods("POST")

	// Protected API routes
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.JWTAuthMiddleware)

	// REST: Article routes
	api.HandleFunc("/articles", handlers.GetArticles(dbConn)).Methods("GET")
	api.HandleFunc("/articles", handlers.SaveArticle(asynqClient, dbConn)).Methods("POST")

	// GraphQL handler
	gqlServer := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{
			Resolvers: &graph.Resolver{
				DB:    dbConn,
				Asynq: asynqClient,
			},
		}),
	)

	// GraphQL routes
	r.Handle("/query", middleware.JWTAuthMiddleware(gqlServer)) // secured GraphQL endpoint
	r.Handle("/", playground.Handler("GraphQL", "/query"))      // GraphQL playground UI

	// Port setup
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
