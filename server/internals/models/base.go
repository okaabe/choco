package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Base struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (this *Base) NewBase() error {

	this.ID = uuid.NewV4().String()
	this.CreatedAt = time.Now()
	this.UpdatedAt = time.Now()

	return nil
}
