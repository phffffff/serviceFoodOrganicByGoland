package imageStorage

import (
	"context"
	"go_service_food_organic/common"
	"go_service_food_organic/module/upload/image/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(imageModel.Image{}.GetTableName()).
		Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
