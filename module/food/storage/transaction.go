package foodStorage

import (
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) BeginTransaction() error {
	ts := sql.db.Table(foodModel.Food{}.TableName()).Begin()
	if err := ts.Error; err != nil {
		return common.ErrDB(err)
	}
	sql.db = ts
	return nil
}

func (sql *sqlModel) RollbackTransaction() error {
	if sql.db == nil {
		return nil
	}
	if err := sql.db.Rollback().Error; err != nil && err != gorm.ErrInvalidTransaction {
		return common.ErrDB(err)
	}
	sql.db = nil
	return nil
}

func (sql *sqlModel) CommitTransaction() error {
	if sql.db == nil {
		return nil
	}
	err := sql.db.Commit().Error
	if err != nil && err != gorm.ErrInvalidTransaction {
		return common.ErrDB(err)
	}
	sql.db = nil
	return nil
}
