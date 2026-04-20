package httpapi

import (
	"fmt"
	"net/http"
)

func (conn *Connection) DeleteTask(w http.ResponseWriter, r *http.Request) {
	sqlQuery := `
		DELETE FROM tasksdb
		WHERE id = $1
	`
	idStr := r.URL.Query().Get("id")
	_, err := conn.Conn.Exec(r.Context(), sqlQuery, idStr)
	if err != nil {
		fmt.Println(err)
	}
}
