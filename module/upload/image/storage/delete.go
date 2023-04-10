package imageStorage

import (
	"context"
	"go_service_food_organic/common"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(common.Image{}.GetTableName()).Updates(map[string]interface{}{"id": id}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
