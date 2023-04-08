package userStorage

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

func (sql *sqlModel) DeleteUser(c context.Context, idUser int) error {
	if err := sql.db.Table(userModel.User{}.GetTableName()).
		Where("id = (?)", idUser).
		Updates(map[string]interface{}{"status": 0}).
		Error; err != nil {
		return common.ErrDB(err)
	}
	return nil

}
