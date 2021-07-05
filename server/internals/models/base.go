package models

import "time"

type Base struct {
	ID        string    `json:"id" gorm:"primaryKey;type:uuid;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
