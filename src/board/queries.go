package board

import (
	"github.com/shinjiezumi/echodock/src/models/board"
	"gorm.io/gorm"
)

func GetBoardList(db *gorm.DB) []*board.Board {
	ret := make([]*board.Board, 0)

	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}

	return ret
}

func GetBoardByID(db *gorm.DB, id int) *board.Board {
	var ret board.Board

	if err := db.Preload("Tags").First(&ret, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		panic(err)
	}

	return &ret
}

func GetTags(db *gorm.DB) []*board.Tag {
	ret := make([]*board.Tag, 0)
	if err := db.Order("id ASC").Find(&ret).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			panic(err)
		}
	}
	return ret
}

func SaveBoard(db *gorm.DB, b *board.Board) {
	if err := db.Create(b).Error; err != nil {
		panic(err)
	}
}

func SaveTagRelation(db *gorm.DB, tr *[]board.TagRelation) {
	if err := db.Create(tr).Error; err != nil {
		panic(err)
	}
}
