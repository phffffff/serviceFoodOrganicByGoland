package userStorage

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

func (sql *sqlModel) Update(c context.Context, id int, pass string) error {
	if err := sql.db.
		Table(userModel.UserPasswordUpdate{}.TableName()).
		Where("id = (?)", id).
		Update("password", pass).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
