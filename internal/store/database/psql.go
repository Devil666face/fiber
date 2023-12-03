package database

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (d *Database) NewPsql() error {
	var err error
	d.PsqlConnect = d.getDNS()
	if d.db, err = gorm.Open(postgres.New(postgres.Config{
		DSN: d.PsqlConnect,
	}), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	}); err != nil {
		return err
	}
	return nil
}

func (d *Database) getDNS() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Moscow statement_timeout=0",
		d.config.PostgresHost,
		d.config.PostgresUser,
		d.config.PostgresPassword,
		d.config.PostgresDB,
		d.config.PostgresPort,
	)
}
