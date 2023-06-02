package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Todo struct {
	ID       uuid.UUID `json:"id" gorm:"primaryKey"`
	Title    string    `json:"title"`
	Note     string    `json:"note"`
	Complete bool      `json:"complete"`
	Owner    string    `json:"-" gorm:"foreignKey"`
}

func (new *Todo) BeforeCreate(tx *gorm.DB) (err error) {
	new.ID = uuid.New()
	return
}
