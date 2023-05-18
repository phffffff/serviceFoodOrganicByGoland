package commentStorage

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
)

func (sql *sqlModel) ListDataWithCondition(c context.Context, filter *commentModel.Filter, paging *common.Paging) ([]commentModel.Comment, error) {
	var list []commentModel.Comment
	db := sql.db.Table(commentModel.Comment{}.TableName())
	if err := db.Error; err != nil {
		return nil, common.ErrDB(err)
	}
	//Cần khai báo những food cho phép hiện
	if len(filter.Status) > 0 {
		db = db.Where("status in (?)", filter.Status)
	}

	if filter.ProfileId > 0 {
		db = db.Where("profile_id = (?)", filter.ProfileId)
	}

	if filter.NewId > 0 {
		db = db.Where("new_id = (?)", filter.NewId)
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
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
		lastData.Mask(false)
		paging.NextCursor = lastData.FakeId.String()
	}

	return list, nil

}
