package addressStorage

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
)

func (sql *sqlModel) Delete(c context.Context, id int) error {
	if err := sql.db.Table(addressModel.Address{}.TableName()).Where("id = (?)", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
