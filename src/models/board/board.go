package board

import (
	"time"
)

type Board struct {
	ID        int `gorm:"primaryKey"`
	Title     string
	Body      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time

	Tags []*Tag `gorm:"many2many:board_tag_relation"`
}

func (b Board) TableName() string {
	return "boards"
}
