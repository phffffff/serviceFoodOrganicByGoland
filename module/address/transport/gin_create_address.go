package addressTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	addressBusiness "go_service_food_organic/module/address/business"
	addressModel "go_service_food_organic/module/address/model"
	addressStorage "go_service_food_organic/module/address/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinCreateAddress(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := addressStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		biz := addressBusiness.NewCreateAddressBiz(store, storeProfile, req)

		var data addressModel.AddressCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		profileUID, err := common.FromBase58(data.ProfileFakeId)
		if err != nil {
			panic(err)
		}
		data.ProfileId = int(profileUID.GetLocalID())

		provinceUID, err := common.FromBase58(data.ProvinceFakeId)
		if err != nil {
			panic(err)
		}
		data.ProvinceId = int(provinceUID.GetLocalID())

		if err := biz.CreateAddress(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
