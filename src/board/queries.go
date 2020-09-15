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
