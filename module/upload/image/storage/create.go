package imageStorage

import (
	"context"
	"go_service_food_organic/common"
)

func (sql *sqlModel) Create(c context.Context, data *common.Image) error {
	if err := sql.db.Table(common.Image{}.GetTableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
