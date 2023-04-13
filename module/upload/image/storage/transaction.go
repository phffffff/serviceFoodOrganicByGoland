package imageStorage

import (
	"go_service_food_organic/common"
	"go_service_food_organic/module/upload/image/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) BeginTransaction() (*gorm.DB, error) {
	db := sql.db.Table(imageModel.Image{}.GetTableName()).Begin()
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return db, nil
}

func (sql *sqlModel) RollbackTransaction(db *gorm.DB) error {
	if err := db.Rollback().Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) CommitTransaction(db *gorm.DB) error {
	err := db.Commit().Error
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
