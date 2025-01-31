package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (b *BaseModel) BeforeCreate(tx *gorm.DB) error {
	uuid := uuid.New()
	tx.Statement.SetColumn("ID", uuid)
	return nil
}
