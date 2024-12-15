package main

import (
	"database/sql"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/vihaan404/metaverse-go/internal/database"
)

type config struct {
	port  int
	query *database.Queries
}

type application struct {
	config config
}

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	dbQueries := database.New(db)
	app := application{
		config: config{
			port:  4000,
			query: dbQueries,
		},
	}

	err = app.serve()
	if err != nil {
		panic(err)
	}
}
