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

	Tags []*Tag `gorm:"many2many:board_tag_relation"`
}

func (b Comment) TableName() string {
	return "comments"
}
