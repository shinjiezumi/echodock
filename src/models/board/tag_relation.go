package board

type TagRelation struct {
	ID      uint `gorm:"primaryKey"`
	BoardID uint
	TagID   uint
}

func (t TagRelation) TableName() string {
	return "board_tag_relation"
}
