package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sql *sqlModel) Create(c context.Context, data *profileModel.ProfileRegister) error {
	if err := sql.db.Table(profileModel.ProfileRegister{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
