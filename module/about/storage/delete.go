package aboutStorage

import (
	"context"
	"go_service_food_organic/common"
	aboutModel "go_service_food_organic/module/about/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(aboutModel.About{}.TableName()).Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
