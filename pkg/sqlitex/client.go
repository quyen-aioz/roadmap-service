package sqlitex

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3" // sqlite3 driver
)

var instance *sql.DB

func InitDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite: %w", err)
	}

	if err := createTable(db); err != nil {
		return nil, fmt.Errorf("failed to create table: %w", err)
	}

	instance = db
	return db, nil
}

func createTable(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS roadmap (
			id TEXT PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			status TEXT NOT NULL,
			group_id TEXT,
			start_date DATETIME NOT NULL,
			end_date DATETIME NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			deleted_at DATETIME
		);
	`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	return nil
}

func Get() (*sql.DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("sqlite not initialized")
	}

	return instance, nil
}
