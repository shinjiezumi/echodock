package board

type TagRelation struct {
	ID      int `gorm:"primaryKey"`
	BoardID int
	TagID   int
}

func (t TagRelation) TableName() string {
	return "board_tag_relation"
}
