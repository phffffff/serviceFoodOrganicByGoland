package commentStorage

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
)

func (sql *sqlModel) Create(c context.Context, data *commentModel.CommentCrt) error {
	if err := sql.db.Table(commentModel.CommentCrt{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
