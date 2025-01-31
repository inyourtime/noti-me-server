package domain

import (
	"github.com/google/uuid"
)

type User struct {
	BaseModel
	Email     string  `gorm:"uniqueIndex;not null"`
	FirstName *string `gorm:"type:varchar(255)"`
	LastName  *string `gorm:"type:varchar(255)"`
	Picture   *string

	Providers []UserProvider `gorm:"foreignKey:UserID"`
}

type UserProvider struct {
	BaseModel
	UserID       uuid.UUID `gorm:"index;type:uuid;not null"`
	Provider     string    `gorm:"type:varchar(50);not null"`
	ProviderID   *string   `gorm:"type:varchar(255)"`
	PasswordHash *string   `gorm:"type:varchar(255)"`
}
