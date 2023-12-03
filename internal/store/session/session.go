package session

import (
	"time"

	"github.com/Devil666face/fiber/internal/config"
	"github.com/Devil666face/fiber/internal/store/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

const (
	table = "storage"
)

type Store struct {
	store         *session.Store
	storage       fiber.Storage
	config        *config.Config
	sqliteConnect string
	psqlConnect   string
}

func New(_config *config.Config, _database *database.Database) *Store {
	s := Store{
		config:        _config,
		sqliteConnect: _database.SqliteConnect,
		psqlConnect:   _database.PsqlConnect,
	}
	s.storage = s.newStorage()
	s.store = s.newStore()
	return &s
}

func (s *Store) Store() *session.Store {
	return s.store
}

func (s *Store) Storage() fiber.Storage {
	return s.storage
}

func (s *Store) newStorage() fiber.Storage {
	if s.config.PostgresUse {
		return s.NewPsqlStorage()
	}
	return s.NewSqliteStorage()
}

func (s *Store) newStore() *session.Store {
	return session.New(session.Config{
		// CookieSecure: true,
		CookieHTTPOnly: true,
		Expiration:     time.Hour * 5,
		Storage:        s.storage,
	})
}
