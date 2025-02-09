package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"unique;not null;type:varchar(50)"`
	Email        string    `gorm:"unique;not null"`
	PasswordHash string    `gorm:"not null"`
	Role         string    `gorm:"type:varchar(50);not null;check:role IN ('customer', 'admin')"`
	CreatedAt    time.Time `gorm:"default:now()"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
