package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (conn *Connection) NewTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, "неверно заполненные даныне", http.StatusBadRequest)
	}
	fmt.Println("--------------------------------------------------------")
	fmt.Println(task.Title, task.Status, task.Description, task.Created_at)

	sqlQuery := `
	INSERT INTO tasksDB(title,description,status,created_at)
	VALUES($1,$2,$3,$4)
	`
	_, error := conn.Conn.Exec(r.Context(), sqlQuery, task.Title, task.Description, task.Status, time.Now())
	if error != nil {
		http.Error(w, "неверно заполненные даныне", http.StatusBadRequest)
		fmt.Println(error)
	}
}
