package database

import (
    "database/sql"
    "log"

    _ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
    var err error
    DB, err = sql.Open("sqlite3", "ideas.db")
    if err != nil {
        log.Fatal(err)
    }
    createTable()
}

func createTable() {
    sqlStmt := `CREATE TABLE IF NOT EXISTS ideas (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT,
        content TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
    _, err := DB.Exec(sqlStmt)
    if err != nil {
        log.Printf("テーブル作成に失敗しました: %q", err)
    }
}
