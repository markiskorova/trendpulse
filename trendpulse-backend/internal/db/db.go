package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    dsn := os.Getenv("DATABASE_URL")
    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    if err = DB.Ping(); err != nil {
        log.Fatalf("DB unreachable: %v", err)
    }
    fmt.Println("Database connected.")
}
