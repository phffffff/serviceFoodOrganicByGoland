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

	db := sqlModel.db.Table(profileModel.Profile{}.GetTableName())

	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if filter.Status >= 0 {
		db = db.Where("status = (?)", filter.Status)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if cursor := paging.FakeCursor; cursor != "" {

		uid, err := common.FromBase58(cursor)

		id := int(uid.GetLocalID())
		if err != nil {
			return nil, common.ErrInternal(err)
		}

		if err := db.Where("id < (?)", id).Error; err != nil {
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
