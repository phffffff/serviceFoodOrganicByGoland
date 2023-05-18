package newStorage

import (
	"context"
	"go_service_food_organic/common"
	newModel "go_service_food_organic/module/new/model"
)

func (sql *sqlModel) Update(c context.Context, data *newModel.New, id int) error {
	if err := sql.db.Table(newModel.New{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
