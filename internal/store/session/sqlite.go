package session

import (
	"github.com/gofiber/storage/sqlite3"
)

func (s *Store) NewSqliteStorage() *sqlite3.Storage {
	return sqlite3.New(sqlite3.Config{
		Database: s.sqliteConnect,
		Table:    table,
		Reset:    false,
	})
}
