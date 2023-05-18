package commentStorage

import (
	"context"
	"go_service_food_organic/common"
	commentModel "go_service_food_organic/module/comment/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*commentModel.Comment, error) {
	var cmt commentModel.Comment
	if err := sql.db.Table(commentModel.Comment{}.TableName()).Where(cond).First(&cmt).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(commentModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &cmt, nil
}
