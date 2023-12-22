package handlers

import (
	"html/template"
	jwthandler "inline/jwt_handler"
	"net/http"
)

func tmplShort(fileString string) *template.Template {
	tmpl := template.Must(template.ParseFiles(fileString))
	return tmpl
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// c := context.Background()
	// client := database.DbAccess(c)

	queueNumber := 0

	jwthandler.CheckCookies(w, r)

	tmpl := tmplShort("./templ/index.html")
	tmpl.Execute(w, map[string]interface{}{"QueueNumber": queueNumber})
}
