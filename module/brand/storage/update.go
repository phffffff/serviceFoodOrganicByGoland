package brandStorage

import (
	"context"
	"go_service_food_organic/common"
	brandModel "go_service_food_organic/module/brand/model"
)

func (sql *sqlModel) Update(c context.Context, data *brandModel.Brand, id int) error {
	if err := sql.db.Table(brandModel.Brand{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
