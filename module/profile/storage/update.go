package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sql *sqlModel) Update(c context.Context, id int, data *profileModel.Profile) error {
	if err := sql.db.Where("id = (?)", id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
