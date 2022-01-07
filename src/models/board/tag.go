package board

import (
	"time"
)

// Tag タグ
type Tag struct {
	ID        int       `gorm:"primaryKey"` // ID
	Name      string    // タグ名
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

// TableName はテーブル名を返します
func (t Tag) TableName() string {
	return "tags"
}
