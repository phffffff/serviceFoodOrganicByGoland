package addressStorage

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
)

func (sql *sqlModel) Create(c context.Context, data *addressModel.AddressCreate) error {
	if err := sql.db.Table(addressModel.AddressCreate{}.TableName()).Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
