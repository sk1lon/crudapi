package main

import (
	"fmt"
	"os"
)

// import (
// 	"context"
// 	"fmt"
// 	httpapi "mogoose/httpAPI"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/jackc/pgx/v5"
// )

// func main() {
// 	ctx := context.Background()
// 	conn, err := pgx.Connect(ctx, "postgres://postgres:skj09@localhost:5432/postgres")
// 	if err != nil {
// 		panic(err)
// 	}
// 	db := &httpapi.Connection{
// 		Conn: conn,
// 	}
// 	httpapi.NewDB(conn, ctx)
// 	router := mux.NewRouter()
// 	fmt.Println("LOX")
// 	router.Path("/tasks").Methods("POST").HandlerFunc(db.NewTask)
// 	router.HandleFunc("/tasks", db.FullTasks).Methods("GET")
// 	// router.Path("/tasks").Methods("GET").HandlerFunc(db.OneTask)
// 	// router.HandleFunc("/tasks", db.DeleteTask).Methods("DELETE")
// 	if err = http.ListenAndServe(":9091", router); err != nil {
// 		fmt.Println(err)
// 	}

// }

func main() {
	fmt.Println("GOIDA")
	ld := os.Getenv("CONN_STRING")
	fmt.Println(ld)
}
