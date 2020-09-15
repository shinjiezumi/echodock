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

func TableName() string {
	return "boards"
}
