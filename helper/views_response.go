package helper

import (
	"net/http"
	"text/template"
)

func Views(w http.ResponseWriter, viewsContent string, data map[string]any) {

	templates := template.Must(template.ParseFiles(
		"views/layout/index.html",
		"views/layout/navbar.html",
		"views/"+viewsContent,
	))

	if data == nil {
		data = map[string]any{
			"Title": "Todo List",
		}
	}

	err := templates.ExecuteTemplate(w, "index.html", &data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
