package imageFoodStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/image_food/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.
		Table(imageFoodModel.ImageFood{}.TableName()).
		Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
