package database

import (
	"database/sql"
	"time"
	
	_ "github.com/mattn/go-sqlite3"
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

func NewSQLiteDatabase(filePath string) *SQLiteDatabase {
	db, err := sql.Open("sqlite3", filePath)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS logs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			message TEXT,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		panic(err)
	}

	return &SQLiteDatabase{db: db}
}

// func (d *SQLiteDatabase) StoreLog(message string) (Log, error) {
// 	res, err := d.db.Exec("INSERT INTO logs (message) VALUES (?)", message)
// 	if err != nil {
// 		return Log{}, err
// 	}

// 	id, _ := res.LastInsertId()
// 	var logEntry Log
// 	err = d.db.QueryRow(`
// 		SELECT id, message, timestamp 
// 		FROM logs 
// 		WHERE id = ?
// 	`, id).Scan(&logEntry.ID, &logEntry.Message, &logEntry.Timestamp)

// 	return logEntry, err
// }

func (d *SQLiteDatabase) GetLogs() ([]Log, error) {
	rows, err := d.db.Query("SELECT id, message, timestamp FROM logs ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.ID, &log.Message, &log.Timestamp)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	return logs, nil
}

func (d *SQLiteDatabase) Close() error {
	return d.db.Close()
}