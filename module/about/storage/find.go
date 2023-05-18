package aboutStorage

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*aboutModel.About, error) {
	var about aboutModel.About
	if err := sql.db.Table(aboutModel.AboutUpdate{}.TableName()).Where(cond).First(&about).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(aboutModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &about, nil
}
