package board

import (
	"time"
)

type Comment struct {
	ID        int `gorm:"primaryKey"`
	BoardID   int
	Comment   string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b Comment) TableName() string {
	return "comments"
}
