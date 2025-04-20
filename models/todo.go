package models

import "golang-todolist/database"

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"status"`
	Increment   int    `json:"increment"`
}

func GetAllTodos() ([]Todo, error) {
	rows, err := database.DB.Query("SELECT id, title, done, description FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	inc := 1
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.Description); err != nil {
			return nil, err
		}
		t.Increment = inc
		inc++
		todos = append(todos, t)
	}
	return todos, nil
}

func InsertTodo(t *Todo) (int, error) {
	SQL := "INSERT INTO todos (title, description) VALUES ($1, $2)"
	_, err := database.DB.Exec(SQL, t.Title, t.Description)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func DeleteTodo(id string) (int, error) {
	SQL := "DELETE FROM todos WHERE id = $1"
	_, err := database.DB.Exec(SQL, id)
	if err != nil {
		return 0, err
	}

	return 1, nil
}

func GetTodoById(id string) (*Todo, error) {
	SQL := "SELECT id, title, done, description FROM todos WHERE id = $1"
	row := database.DB.QueryRow(SQL, id)

	var t Todo
	if err := row.Scan(&t.ID, &t.Title, &t.Done, &t.Description); err != nil {
		return &t, err
	}

	return &t, nil
}

func UpdateTodo(t *Todo) (int, error) {
	SQL := "UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4"
	_, err := database.DB.Exec(SQL, t.Title, t.Description, t.Done, t.ID)
	if err != nil {
		return 0, err
	}

	return 1, nil
}
