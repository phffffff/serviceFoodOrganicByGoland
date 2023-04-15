package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithConditon(
	c context.Context,
	cond map[string]interface{},
	morekeys ...string) (*profileModel.Profile, error) {
	var data profileModel.Profile
	if err := sql.db.Table(profileModel.Profile{}.TableName()).Where(cond).First(&data).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(profileModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
