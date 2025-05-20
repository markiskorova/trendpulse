package main

import (
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "trendpulse/internal/handlers"
    "trendpulse/internal/middleware"
    "trendpulse/internal/db"
)

func main() {
    db.InitDB()

    r := mux.NewRouter()
    r.HandleFunc("/health", handlers.HealthCheck).Methods("GET")
    r.HandleFunc("/register", handlers.Register).Methods("POST")
    r.HandleFunc("/login", handlers.Login).Methods("POST")

    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.JWTAuthMiddleware)
    api.HandleFunc("/events", handlers.SubmitEvent).Methods("POST")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    log.Printf("Server starting on port %s...", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}
