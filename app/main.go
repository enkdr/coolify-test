package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Item struct {
	ID          int
	Name        string
	Description string
}

func main() {
	connStr := "postgres://myuser:mypassword@postgres:5432/mydatabase?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT id, name, description FROM items")
		if err != nil {
			http.Error(w, "Database query failed", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		items := []Item{}
		for rows.Next() {
			var item Item
			if err := rows.Scan(&item.ID, &item.Name, &item.Description); err != nil {
				http.Error(w, "Failed to scan item", http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}

		for _, item := range items {
			fmt.Fprintf(w, "ID: %d, Name: %s, Description: %s\n", item.ID, item.Name, item.Description)
		}
	})

	log.Fatal(http.ListenAndServe(":3000", nil))
}
