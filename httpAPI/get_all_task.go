package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func (conn *Connection) FullTasks(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
		SELECT id, title, description, status, created_at 
		FROM tasksdb
		LIMIT $1 OFFSET $2
	`

	page := r.URL.Query().Get("page")
	pages, _ := strconv.Atoi(page)
	limit := r.URL.Query().Get("limit")
	limits, _ := strconv.Atoi(limit)
	offset := (pages - 1) * limits
	rows, err := conn.Conn.Query(r.Context(), sqlQuery, limits, offset)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.Status, &task.Created_at); err != nil {
			fmt.Println(err)
		}
		fmt.Println(task)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}
