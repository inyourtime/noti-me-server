package repository

import (
	"github.com/inyourtime/noti-me-server/internal/core/port"
	"gorm.io/gorm"
)

type atomicRepository struct {
	db *gorm.DB
}

type repository struct {
	db *gorm.DB
}

func NewAtomicRepository(db *gorm.DB) port.AtomicRepository {
	return &atomicRepository{db: db}
}

func (r *atomicRepository) Transaction(callback port.AtomicRepoositoryCallback) error {
	err := r.db.Transaction(func(tx *gorm.DB) error {
		return callback(create(tx))
	})

	return err
}

func create(db *gorm.DB) port.Repository {
	return &repository{db: db}
}

func (r *repository) UserRepository() port.UserRepository {
	return NewUserRepository(r.db)
}
