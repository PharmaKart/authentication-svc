package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Customer struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;unique;not null"`
	FirstName   string    `gorm:"not null"`
	LastName    string    `gorm:"not null"`
	Phone       *string
	DateOfBirth *time.Time
	StreetLine1 string `gorm:"not null"`
	StreetLine2 *string
	City        string    `gorm:"not null"`
	Province    string    `gorm:"not null"`
	PostalCode  string    `gorm:"not null"`
	Country     string    `gorm:"not null;default:Canada"`
	CreatedAt   time.Time `gorm:"type:timestamptz;default:now()"`
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}
