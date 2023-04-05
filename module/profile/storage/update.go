package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sql *sqlModel) Update(c context.Context, data *profileModel.Profile) error {
	if err := sql.db.Where("id = (?)", data.Id).Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
