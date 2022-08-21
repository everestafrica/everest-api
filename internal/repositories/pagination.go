package repositories

import "gorm.io/gorm"

func paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {

		offset := (page - 1) * size
		return db.Offset(offset).Limit(size)
	}
}
