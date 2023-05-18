package commentStorage

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(commentModel.Comment{}.TableName()).Where("id = (?)", id).
		Delete(nil).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
