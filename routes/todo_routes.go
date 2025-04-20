package routes

import (
	"golang-todolist/controllers"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", controllers.Home)
	router.GET("/create-todo", controllers.CreateTodo)
	router.POST("/insert-todo", controllers.InsertTodo)
	router.DELETE("/delete-todo/:id", controllers.DeleteTodo)
	router.GET("/edit-todo/:id", controllers.EditTodo)
	router.POST("/update-todo/:id", controllers.UpdateTodo)

	return router
}
