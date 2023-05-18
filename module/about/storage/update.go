package aboutStorage

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

func (sql *sqlModel) Update(c context.Context, data *aboutModel.About, id int) error {
	if err := sql.db.Table(aboutModel.About{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
