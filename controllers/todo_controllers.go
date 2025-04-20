package controllers

import (
	"golang-todolist/helper"
	"golang-todolist/models"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	todo, err := models.GetAllTodos()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := map[string]any{
		"Title": "Todo List",
		"Todos": todo,
		"Inc": func(a int) int {
			return a + 1
		},
	}

	helper.Views(w, "page/home.html", data)
}

func CreateTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	helper.Views(w, "page/create.html", nil)
}

func InsertTodo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	todo := models.Todo{
		Title:       title,
		Description: description,
	}
	_, err := models.InsertTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	_, err := models.DeleteTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func EditTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	todo, err := models.GetTodoById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := map[string]any{
		"Title": "Edit Todo",
		"Todo":  todo,
	}
	helper.Views(w, "page/edit.html", data)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id, _ := strconv.Atoi(p.ByName("id"))
	r.ParseForm()
	title := r.FormValue("title")
	description := r.FormValue("description")
	done := r.FormValue("done") == "1"

	todo := models.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Done:        done,
	}

	_, err := models.UpdateTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
