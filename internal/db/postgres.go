package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
	// DB driver
	_ "github.com/lib/pq"
)

func NewDB(dbHost, dbPort, dbName, dbUser, dbPassword string ) (newDB *sqlx.DB, err error) {
	dbUri := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHost, dbPort)
	log.Println("successfully connection to DB")
	log.Println("dbUri :", dbUri)
	db, err := sql.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(60)
	db.SetMaxOpenConns(60)
	db.SetConnMaxLifetime(5 * time.Minute)

	newDB = sqlx.NewDb(db, "postgres")
	return newDB, nil
}
