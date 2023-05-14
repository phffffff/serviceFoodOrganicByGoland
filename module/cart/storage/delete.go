package cartStorage

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/cart/model"
)

func (sql *sqlModel) Delete(c context.Context, userId int) error {
	if err := sql.db.Table(cartModel.Cart{}.TableName()).Where("user_id = (?)", userId).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
