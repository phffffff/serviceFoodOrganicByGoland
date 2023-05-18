package newTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	newBusiness "go_service_food_organic/module/new/business"
	newModel "go_service_food_organic/module/new/model"
	newStorage "go_service_food_organic/module/new/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinCreateNew(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		storeProfile := profileStorage.NewSqlModel(db)
		store := newStorage.NewSqlModel(db)

		biz := newBusiness.NewCreateNewBiz(store, storeProfile, req)

		var data newModel.NewCrt
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		if err := biz.CreateNew(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
