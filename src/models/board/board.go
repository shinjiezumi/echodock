package board

import (
	"time"
)

// Board 掲示板投稿
type Board struct {
	ID        int       `gorm:"primaryKey"` // ID
	Title     string    // タイトル
	Body      string    // 本文
	Name      string    // 投稿者名
	CreatedAt time.Time // 作成日時
	UpdatedAt time.Time // 更新日時

	Tags     []*Tag     `gorm:"many2many:board_tag_relation"` // タグ
	Comments []*Comment // コメント
}

// TableName はテーブル名を返します
func (b Board) TableName() string {
	return "boards"
}
