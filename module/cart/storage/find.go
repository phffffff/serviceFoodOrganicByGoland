package cartStorage

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/cart/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*cartModel.Cart, error) {
	var cart cartModel.Cart
	if err := sql.db.Table(cartModel.Cart{}.TableName()).Where(cond).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(cartModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &cart, nil
}
