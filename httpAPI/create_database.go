package httpapi

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewDB(conn *pgx.Conn, ctx context.Context) {
	sqlBD := `
		CREATE TABLE IF NOT EXISTS tasksDB (
		id SERIAL PRIMARY KEY,
		title VARCHAR(200) NOT NULL UNIQUE,
		description VARCHAR(1000) NOT NULL,
		status VARCHAR(10) NOT NULL,
		created_at TIMESTAMP NOT NULL)
	`
	_, err := conn.Exec(ctx, sqlBD)
	if err != nil {
		panic(err)
	}
}
