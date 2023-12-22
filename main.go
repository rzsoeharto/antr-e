package main

import (
	"context"
	"fmt"
	logger "inline/Logger"
	"inline/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/lpernett/godotenv"
)

type Handler func(w http.ResponseWriter, c context.Context, r *http.Request) error

func init() {
	err := godotenv.Load()
	logger.Init("./AppLog.log")

	if err != nil {
		logger.Logger.Panicln(err)
		return
	}
}

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", handlers.HomeHandler)

	fmt.Println("Server online")

	http.ListenAndServe(":8080", r)
}
