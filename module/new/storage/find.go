package newStorage

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*newModel.New, error) {
	var new newModel.New
	if err := sql.db.Table(newModel.New{}.TableName()).Where(cond).First(&new).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(newModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &new, nil
}
