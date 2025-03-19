package database

import (
	"database/sql"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	
)

type Log struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

type Database interface {
	StoreLog(message string) (Log, error)
	Close() error
}

type SQLiteDatabase struct {
	db *sql.DB
}


func NewSQLiteDatabase(filePath string) (*SQLiteDatabase, error) {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening database: %v", err)
	}

	

	query := `CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY AUTOINCREMENT, message TEXT, timestamp DATETIME DEFAULT CURRENT_TIMESTAMP);`
	_, err = db.Exec(query)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("Error creating table: %v", err)
	}

	return &SQLiteDatabase{db: db}, nil
}



func (s *SQLiteDatabase) StoreLog(message string) (Log, error) {
	// Check if the database connection is initialized
	if s == nil || s.db == nil {
		return Log{}, fmt.Errorf("database connection is not initialized")
	}

	res, err := s.db.Exec("INSERT INTO logs (message) VALUES (?)", message)
	if err != nil {
		return Log{}, err
	}

	id, _ := res.LastInsertId()
	var logEntry Log
	err = s.db.QueryRow(`
		SELECT id, message, timestamp 
		FROM logs 
		WHERE id = ?
	`, id).Scan(&logEntry.ID, &logEntry.Message, &logEntry.Timestamp)

	if err != nil {
		return Log{}, fmt.Errorf("Error fetching log entry: %v", err)
	}

	return logEntry, nil
}


func (d *SQLiteDatabase) Close() error {
	return d.db.Close()
}