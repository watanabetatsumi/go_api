package main

import (
	"database/sql"
	"log"
	"net/http"
)

type TodoItem struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func main() {
	db, err := sql.Open("mysql", "root:Wtatasumi0317@tcp(127.0.0.1.3306)/todoapp")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/todo", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		}
	})
}
