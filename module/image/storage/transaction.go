package imageStorage

import (
	"go_service_food_organic/common"
	imageModel "go_service_food_organic/module/image/model"
)

func (sql *sqlModel) BeginTransaction() error {
	db := sql.db.Table(imageModel.Image{}.TableName()).Begin()
	if err := db.Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) RollbackTransaction() error {
	if err := sql.db.Rollback().Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}

func (sql *sqlModel) CommitTransaction() error {
	err := sql.db.Commit().Error
	if err != nil {
		return common.ErrDB(err)
	}
	return nil
}
