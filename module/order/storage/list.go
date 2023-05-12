package orderStorage

import (
	"context"
	"go_service_food_organic/common"
	orderModel "go_service_food_organic/module/order/model"
)

func (sqlModel *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *orderModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]orderModel.Order, error) {

	db := sqlModel.db.Table(orderModel.Order{}.TableName())

	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(filter.Status) >= 0 {
		db = db.Where("status = (?)", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range moreKeys {
		db.Preload(item)
	}

	if cursor := paging.FakeCursor; cursor != "" {

		uid, err := common.FromBase58(cursor)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		id := int(uid.GetLocalID())

		db = db.Where("id < (?)", id)
		if err := db.Error; err != nil {
			return nil, common.ErrInternal(err)
		}
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}
	var list []orderModel.Order

	if err := db.Limit(paging.Limit).Order("id DESC").Find(&list).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(list) > 0 {
		lastData := list[len(list)-1]
		lastData.Mark(false)
		paging.NextCursor = lastData.FakeId.String()
	}

	//mark
	for _, item := range list {
		item.Users.Mark(false)
		for _, od := range item.OrderDetails {
			od.Mark(false)
			od.Foods.Mark(false)
		}
	}

	return list, nil

}
