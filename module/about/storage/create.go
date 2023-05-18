package aboutStorage

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

func (sql *sqlModel) Create(c context.Context, data *aboutModel.About) error {
	if err := sql.db.Table(aboutModel.About{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
