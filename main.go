package main

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("./templ/index.html"))
		queueNumber := rand.Intn(30)

		tmpl.Execute(w, map[string]interface{}{"QueueNumber": queueNumber})
	})

	fmt.Println("Server online")
	http.ListenAndServe(":8080", r)
}
