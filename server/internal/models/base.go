package models

import (
	"time"

	uuid "github.com/gofrs/uuid"
)

//What every data on the applicatiom must have
type Base struct {
	ID        string    `json:"id" gorm:"primaryKey; type:uuid;"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (base *Base) id() error {
	id, err := uuid.NewV4()

	if err != nil {
		return err
	}

	base.ID = id.String()

	return nil
}
