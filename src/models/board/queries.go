package board

import (
	"gorm.io/gorm"
)

// GetBoardList は掲示板一覧を取得します
func GetBoardList(db *gorm.DB) []*Board {
	ret := make([]*Board, 0)

	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}

	return ret
}

// GetBoardByID はIDで掲示板を取得します
func GetBoardByID(db *gorm.DB, id int) *Board {
	var ret Board

	if err := db.
		Preload("Tags").
		Preload("Comments").
		First(&ret, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &ret
}

// GetTags はタグ一覧を取得します
func GetTags(db *gorm.DB) []*Tag {
	ret := make([]*Tag, 0)
	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}
	return ret
}

// SaveBoard は掲示板を保存します
func SaveBoard(db *gorm.DB, b *Board) {
	var current Board
	if err := db.First(&current, b.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			if err := db.Create(b).Error; err != nil {
				panic(err)
			}
			return
		}
	}

	b.ID = current.ID
	if err := db.Save(b).Error; err != nil {
		panic(err)
	}
}

// DeleteBoard は投稿を削除します
func DeleteBoard(db *gorm.DB, boardID int) {
	if err := db.Delete(Board{}, boardID).Error; err != nil {
		panic(err)
	}
}

// SaveTagRelation はタグリレーションを保存します
func SaveTagRelation(db *gorm.DB, tr *[]TagRelation) {
	if err := db.Create(tr).Error; err != nil {
		panic(err)
	}
}

// DeleteTagRelation はタグリレーションを削除します
func DeleteTagRelation(db *gorm.DB, boardID int) {
	if err := db.Where("board_id = ?", boardID).Delete(TagRelation{}).Error; err != nil {
		panic(err)
	}
}

// GetComment は掲示板IDとコメントIDでコメントを取得します
func GetComment(db *gorm.DB, boardID, commentID int) *Comment {
	var ret Comment

	if err := db.
		Where("board_id = ?", boardID).
		Where("id = ?", commentID).
		First(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &ret
}

// SaveComment はコメントを保存します
func SaveComment(db *gorm.DB, c *Comment) {
	if err := db.Create(c).Error; err != nil {
		panic(err)
	}
}

// DeleteComment はコメントを削除します
func DeleteComment(db *gorm.DB, c *Comment) {
	if err := db.
		Where("id = ?", c.ID).
		Delete(Comment{}).Error; err != nil {
		panic(err)
	}
}
