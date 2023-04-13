package imageStorage

import (
	"context"
	"go_service_food_organic/common"
	imageModel "go_service_food_organic/module/upload/image/model"
)

func (sqlModel *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *imageModel.Filter,
	paging *common.Paging) ([]imageModel.Image, error) {

	db := sqlModel.db.Table(imageModel.Image{}.GetTableName())

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
	var list []imageModel.Image

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
