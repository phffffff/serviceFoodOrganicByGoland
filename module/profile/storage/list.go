package profileStorage

import (
	"context"
	"go_service_food_organic/common"
	profileModel "go_service_food_organic/module/profile/model"
)

func (sql *sqlModel) ListDataWithFilter(
	c context.Context,
	filter *profileModel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]profileModel.Profile, error) {

	var list []profileModel.Profile

	db := sql.db.Table(profileModel.Profile{}.TableName())

	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if filter.Status >= 0 {
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
		list[idx].Image.Mark(false)
	}

	return list, nil

}
