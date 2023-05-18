package addressTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	addressBusiness "go_service_food_organic/module/address/business"
	addressRepo "go_service_food_organic/module/address/repository"
	addressStorage "go_service_food_organic/module/address/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinDeleteAddress(appCtx appContext.AppContext) gin.HandlerFunc {
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
		repo := addressRepo.NewDeleteAddressRepo(store, storeProfile, req)
		biz := addressBusiness.NewDeleteAddressBiz(repo)

		if err := biz.DeleteAddress(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
