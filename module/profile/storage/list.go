package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sqlModel *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *profileModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]profileModel.Profile, error) {

	db := sqlModel.db

	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if filter.Status >= 0 {
		db = db.Where("status = (?)", filter.Status)
	}

	if err := db.Table(profileModel.Profile{}.GetTableName()).Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	//db = db.Preload("Image")
	//for _, item := range moreKeys {
	//	db = db.Preload(item)
	//}

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
	var list []profileModel.Profile

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
