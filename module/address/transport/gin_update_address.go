package addressTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	addressBusiness "go_service_food_organic/module/address/business"
	addressModel "go_service_food_organic/module/address/model"
	addressRepo "go_service_food_organic/module/address/repository"
	addressStorage "go_service_food_organic/module/address/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinUpdateAddress(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := addressStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		repo := addressRepo.NewUpdateAddressRepo(store, storeProfile, req)
		biz := addressBusiness.NewUpdateAddressBiz(repo)

		var data addressModel.AddressUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		profileUID, err := common.FromBase58(data.ProfileFakeId)
		if err != nil {
			panic(err)
		}
		data.ProfileId = int(profileUID.GetLocalID())
		data.ProfileFakeId = ""

		provinceUID, err := common.FromBase58(data.ProvinceFakeId)
		if err != nil {
			panic(err)
		}
		data.ProvinceId = int(provinceUID.GetLocalID())
		data.ProvinceFakeId = ""

		if err := biz.UpdateAddress(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
