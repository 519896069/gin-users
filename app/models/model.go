package models

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID uint `gorm:"primarykey" json:"id"`

	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Db        *gorm.DB       `gorm:"-" json:"-"`
}
