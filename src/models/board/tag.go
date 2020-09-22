package board

import (
	"time"
)

type Tag struct {
	ID        int `gorm:"primaryKey"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (t Tag) TableName() string {
	return "tags"
}
