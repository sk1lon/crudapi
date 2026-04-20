package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (conn *Connection) OneTask(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
		SELECT id,title,description,status,created_at
		FROM tasksdb
		WHERE id = $1
	`
	idStr := r.URL.Query().Get("id")
	var task Task
	err := conn.Conn.QueryRow(r.Context(), sqlQuery, idStr).Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Created_at)
	if err != nil {
		fmt.Println(err)
	}

	json.NewEncoder(w).Encode(task)
	fmt.Println(task)
}
