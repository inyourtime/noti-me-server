package mockdb

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, sqlmock.Sqlmock, *sql.DB) {
	db, mock, _ := sqlmock.New()
	gormDB, _ := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
	return gormDB, mock, db
}
