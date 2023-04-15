package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sql *sqlModel) DeleteProfile(c context.Context, idProfile int) error {
	if err := sql.db.
		Table(profileModel.Profile{}.TableName()).
		Where("id = (?)", idProfile).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
