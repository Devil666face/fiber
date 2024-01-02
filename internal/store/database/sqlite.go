package database

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (d *Database) NewSqlite() error {
	var err error
	if d.db, err = gorm.Open(sqlite.Open(d.SqliteConnect+"?cache=shared&mode=rwc&_busy_timeout=50000"), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().Local() },
		Logger:  logger.Default.LogMode(logger.Info),
	}); err != nil {
		return err
	}
	return nil
}
