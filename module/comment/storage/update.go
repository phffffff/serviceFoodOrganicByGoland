package commentStorage

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
)

func (sql *sqlModel) Update(c context.Context, data *commentModel.Comment, id int) error {
	if err := sql.db.Table(commentModel.CommentUpd{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
