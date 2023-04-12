package main

import (
	"database/sql"
	"encoding/json"
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
		case "GET":
			items, err := getTodoItems(db)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			json.NewEncoder(w).Encode(items)

		case "POST":
			title := r.FormValue("title")
			if title == "" {
				http.Error(w, "title parameter is required", http.StatusBadRequest)
				return
			}

			id, err := addTodoItem(db, title)
			if err != nil {
				http.Error(w, "id parameter is required", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(map[string]int{"id": int(id)})
			// 新しく登録したidを返す
		case "DELETE":
			id := r.FormValue("id")
			if id == "" {
				http.Error(w, "id parameter is required", http.StatusBadRequest)
				return
			}
			json.NewEncoder(w).Encode(map[string]string{"message": "item deleted"})
		default:
			http.Error(w, "unsupported method", http.StatusMethodNotAllowed)
		}
	})
	log.Println("Starting HTTP server on :8080...")
	http.ListenAndServe(":8080", nil)
}

func getTodoItems(db *sql.DB) ([]TodoItem, error) {
	rows, err := db.Query("SELECT id, title FROM items ORDER BY CREATED_at DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []TodoItem
	for rows.Next() {
		var item TodoItem
		err := rows.Scan(&item.ID, &item.Title)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

func addTodoItem(db *sql.DB, title string) (int64, error) {
	result, err := db.Exec("INSERT INTO items (title) VALUES(?)", title)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func deleteTodoItem(db *sql.DB, id string) error {
	_, err := db.Exec("DELETE FROM items WHERE id = ?", id)
	return err
}
