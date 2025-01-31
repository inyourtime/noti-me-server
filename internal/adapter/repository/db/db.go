package db

import (
	"github.com/inyourtime/noti-me-server/config"
	"github.com/inyourtime/noti-me-server/internal/core/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func New(config *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DbUrl), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	pg, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err = pg.Ping(); err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&domain.User{}, &domain.UserProvider{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
