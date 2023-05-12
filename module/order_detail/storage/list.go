package orderDetailStorage

import (
	"context"
	"go_service_food_organic/common"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
)

func (sql *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *orderDetailModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]orderDetailModel.OrderDetail, error) {

	var list []orderDetailModel.OrderDetail

	db := sql.db.Table(orderDetailModel.OrderDetail{}.TableName())

	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(filter.Status) >= 0 {
		db = db.Where("status in (?)", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
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

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if err := db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&list).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(list) > 0 {
		lastData := list[len(list)-1]
		lastData.Mark(false)
		paging.NextCursor = lastData.FakeId.String()
	}

	for idx := range list {
		list[idx].Foods.Mark(false)
		for i := range list[idx].Foods.FoodImages {
			list[idx].Foods.FoodImages[i].Mark(false)
		}
	}

	return list, nil

}
