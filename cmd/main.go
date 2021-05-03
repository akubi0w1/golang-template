package main

import (
	"context"
	"net/http"

	"github.com/akubi0w1/golang-sample/interface/handler"
	"github.com/akubi0w1/golang-sample/interface/persistent/mysql/ent"
	"github.com/akubi0w1/golang-sample/log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	logger := log.New()
	ctx := context.Background()

	dbClient, err := ent.Open("mysql", "worker:password@tcp(db:3306)/main?parseTime=true")
	if err != nil {
		logger.Fatal("failed to open db: %v", err)
	}
	if err := dbClient.Schema.Create(ctx); err != nil {
		logger.Fatal("failed to create schema: %v", err)
	}

	app := handler.NewApp(dbClient)
	mux := app.Routing()

	http.ListenAndServe(":8080", mux)
}
