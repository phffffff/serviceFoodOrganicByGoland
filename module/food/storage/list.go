package foodStorage

import (
	"context"
	"go_service_food_organic/common"
	foodModel "go_service_food_organic/module/food/model"
)

func (sql *sqlModel) ListDataWithCondition(c context.Context, filter *foodModel.Filter, paging *common.Paging, moreKeys ...string) ([]foodModel.Food, error) {
	var list []foodModel.Food
	db := sql.db.Table(foodModel.Food{}.GetTableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	//Cần khai báo những food cho phép hiện
	if len(filter.Status) > 0 {
		db = db.Where("status in (?)", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	for _, item := range moreKeys {
		db = db.Preload(item)
	}

	if cursor := paging.FakeCursor; cursor != "" {
		//id, err := strconv.Atoi(cursor)
		uid, err := common.FromBase58(cursor)
		if err != nil {
			return nil, common.ErrInternal(err)
		}
		id := int(uid.GetLocalID())

		if err != nil {
			return nil, common.ErrInternal(err)
		}
		db = db.Where("id < (?)", id)
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)

	}
	if err := db.Limit(paging.Limit).Order("id DESC").Find(&list).Error; err != nil {
		return nil, common.ErrDB(err)
	}
	if len(list) > 0 {
		lastData := list[len(list)-1]
		lastData.Mark(false)
		paging.NextCursor = lastData.FakeId.String()
	}

	return list, nil

}
