package sql_utils

import (
	"database/sql"
	"strings"
)

// Open is a wrapper for sql.Open which sets the max open connections to 1
// for in memory sqlite3 dbs. In memory sqlite3 creates a new database
// on each connection, so the number of open connections must be limited
// to 1 for thread safety. Otherwise, there is a race condition between
// threads using a cached connection to the original database or opening
// a new connection to a new database.
func Open(driver string, source string) (*sql.DB, error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		return nil, err
	}
	if driver == "sqlite3" && strings.Contains(source, ":memory:") {
		db.SetMaxOpenConns(1)
	}
	return db, nil
}
