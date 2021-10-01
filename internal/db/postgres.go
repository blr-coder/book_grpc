package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
	// DB driver
	_ "github.com/lib/pq"
)

//var dbClient *sqlx.DB

func NewDBClient() (dbClient *sqlx.DB, err error) {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort:= os.Getenv("DB_PORT")

	dbUri := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	dbClient = sqlx.NewDb(db, "postgres")
	return dbClient, nil
}
