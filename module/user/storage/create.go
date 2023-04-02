package userStorage

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

func (sql *sqlModel) Create(c context.Context, data *userModel.UserRegister) error {
	if err := sql.db.Table(userModel.UserLogin{}.GetTableName()).
		Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
