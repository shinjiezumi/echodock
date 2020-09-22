package board

import (
	"gorm.io/gorm"
)

func GetBoardList(db *gorm.DB) []*Board {
	ret := make([]*Board, 0)

	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}

	return ret
}

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

func GetTags(db *gorm.DB) []*Tag {
	ret := make([]*Tag, 0)
	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}
	return ret
}

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

func DeleteBoard(db *gorm.DB, boardID int) {
	if err := db.Delete(Board{}, boardID).Error; err != nil {
		panic(err)
	}
}

func SaveTagRelation(db *gorm.DB, tr *[]TagRelation) {
	if err := db.Create(tr).Error; err != nil {
		panic(err)
	}
}

func DeleteTagRelation(db *gorm.DB, boardID int) {
	if err := db.Where("board_id = ?", boardID).Delete(TagRelation{}).Error; err != nil {
		panic(err)
	}
}

func GetComment(db *gorm.DB, boardID, commentID int) *Comment {
	var ret Comment

	if err := db.
		Where("board_id = ?", boardID).
		Where("id = ?", commentID).
		Find(&ret).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &ret
}

func SaveComment(db *gorm.DB, c *Comment) {
	if err := db.Create(c).Error; err != nil {
		panic(err)
	}
}

func DeleteComment(db *gorm.DB, boardID, commentID int) {
	if err := db.
		Where("board_id = ?", boardID).
		Where("id = ?", commentID).
		Delete(Comment{}).Error; err != nil {
		panic(err)
	}
}
