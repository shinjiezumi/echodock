package board

import (
	"time"
)

// Comment コメント
type Comment struct {
	ID        int       `gorm:"primaryKey"` // ID
	BoardID   int       // 投稿ID
	Comment   string    // コメント
	Name      string    // 投稿者名
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時
}

// TableName はテーブル名を返します
func (b Comment) TableName() string {
	return "comments"
}
