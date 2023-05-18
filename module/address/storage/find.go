package addressStorage

import (
	"context"
	"go_service_food_organic/common"
	addressModel "go_service_food_organic/module/address/model"
	"gorm.io/gorm"
)

func (sql *sqlModel) FindDataWithCondition(c context.Context, cond map[string]interface{}) (*addressModel.Address, error) {
	var address addressModel.Address
	if err := sql.db.Table(addressModel.Address{}.TableName()).Where(cond).First(&address).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.ErrRecordNotFound(addressModel.EntityName, err)
		}
		return nil, common.ErrDB(err)
	}
	return &address, nil
}
