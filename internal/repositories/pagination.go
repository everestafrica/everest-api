package repositories

import (
	"github.com/everestafrica/everest-api/internal/commons/types"
	"gorm.io/gorm"
)

func paginate(p types.Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (p.Page - 1) * p.Size
		return db.Offset(offset).Limit(p.Size)
	}
}
