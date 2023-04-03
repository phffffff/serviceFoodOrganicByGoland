package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sqlModel *sqlModel) Create(c context.Context, data *profileModel.Profile) error {
	if err := sqlModel.db.Table(profileModel.Profile{}.GetTableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
