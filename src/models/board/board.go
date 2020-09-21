package board

import (
	"time"
)

type Board struct {
	ID        uint `gorm:"primaryKey"`
	Title     string
	Body      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b Board) TableName() string {
	return "boards"
}
