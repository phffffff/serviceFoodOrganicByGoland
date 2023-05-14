package cartStorage

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/cart/model"
)

func (sql *sqlModel) Create(c context.Context, data *cartModel.Cart) error {
	if err := sql.db.Table(cartModel.Cart{}.TableName()).Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
