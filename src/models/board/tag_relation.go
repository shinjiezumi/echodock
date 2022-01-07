package board

// TagRelation タグリレーション
// 投稿とタグの関連付け
type TagRelation struct {
	ID      int `gorm:"primaryKey"` // ID
	BoardID int // 投稿ID
	TagID   int // タグID
}

// TableName はテーブル名を返します
func (t TagRelation) TableName() string {
	return "board_tag_relation"
}
