package sqlite

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // I want import init function in code.
	"github.com/yvv4git/users-service/internal/config"
)

// NewSqliteDB is used as constructor for db instance.
func NewSqliteDB(cfg config.DBConfig) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cfg.FileNameDB)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
