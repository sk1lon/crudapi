package httpapi

import (
	"time"

	"github.com/jackc/pgx/v5"
)

type Task struct {
	Id          int
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Created_at  time.Time
}

type Connection struct {
	Conn *pgx.Conn
}
