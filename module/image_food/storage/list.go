package imageFoodStorage

import (
	"context"
	"go_service_food_organic/common"
	imageFoodModel "go_service_food_organic/module/image_food/model"
)

func (sql *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *imageFoodModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]imageFoodModel.ImageFood, error) {
	db := sql.db.Table(imageFoodModel.ImageFood{}.TableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if len(filter.Status) >= 0 {
		db.Where("status in (?)", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	// for _, item := range moreKeys {
	// 	db.Preload(item)
	// }

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
		db.Offset(offset)
	}
	var list []imageFoodModel.ImageFood

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
