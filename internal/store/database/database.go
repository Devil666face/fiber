package database

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/Devil666face/fiber/internal/config"
	"github.com/Devil666face/fiber/internal/models"
	"github.com/Devil666face/fiber/pkg/utils"

	"gorm.io/gorm"
)

type Database struct {
	db            *gorm.DB
	config        *config.Config
	tables        []any
	SqliteConnect string
	PsqlConnect   string
}

func Must(_config *config.Config, _tables []any) *Database {
	d := Database{
		config: _config,
		tables: _tables,
	}
	d.SqliteConnect = d.config.SqliteDB
	d.PsqlConnect = d.getDNS()
	if err := d.connect(); err != nil {
		slog.Error(fmt.Sprintf("Connect database: %s", err))
		//nolint:revive // If database not connect - exit
		os.Exit(1)
	}
	if err := d.migrate(); err != nil {
		slog.Warn(fmt.Sprintf("Migrations not create: %s", err))
	}
	if err := d.createSuperuser(); err != nil {
		slog.Warn(fmt.Sprintf("Superuser not create: %s", err))
	}
	return &d
}

func (d *Database) DB() *gorm.DB {
	return d.db
}

func (d *Database) migrate() error {
	return d.db.AutoMigrate(d.tables...)
}

func (d *Database) connect() error {
	if d.config.PostgresUse {
		return d.NewPsql()
	}
	return d.NewSqlite()
}

func (d *Database) createSuperuser() error {
	hash, err := utils.GenHash(d.config.SuperuserPassword)
	if err != nil {
		return err
	}
	u := models.User{
		Email:    d.config.Superuser,
		Admin:    true,
		Password: hash,
	}
	return u.Create(d.db)
}
