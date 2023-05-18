package addressStorage

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
)

func (sql *sqlModel) Update(c context.Context, data *addressModel.Address, id int) error {
	if err := sql.db.Table(addressModel.Address{}.TableName()).
		Where("id = (?)", id).
		Updates(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
