package cartStorage

import (
	"context"
	"go_service_food_organic/common"
	cartModel "go_service_food_organic/module/cart/model"
)

func (sql *sqlModel) ListDataWithCondition(c context.Context, filter *cartModel.Filter, paging *common.Paging) ([]cartModel.CartLst, error) {
	var list []cartModel.CartLst
	db := sql.db.Table(cartModel.Cart{}.TableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if filter.UserId > 0 {
		db = db.Where("user_id = (?)", filter.UserId)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	offset := (paging.Page - 1) * paging.Limit

	if err := db.Offset(offset).Limit(paging.Limit).Order("food_id DESC").Find(&list).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return list, nil
}
