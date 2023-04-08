package userStorage

import (
	"context"
	"go_service_food_organic/common"
	userModel "go_service_food_organic/module/user/model"
)

func (sql *sqlModel) FindDataWithCondition(
	c context.Context,
	cond map[string]interface{},
	moreKeys ...string) (*userModel.User, error) {

	db := sql.db.Table(userModel.User{}.GetTableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	//for i := range moreKeys {
	//	db = db.Preload(moreKeys[i])
	//}

	var data userModel.User
	if err := db.Where(cond).First(&data).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	return &data, nil
}
