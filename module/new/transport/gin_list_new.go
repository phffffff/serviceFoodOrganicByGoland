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

func GinListNew(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := newStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		biz := newBusiness.NewListNewBiz(store, storeProfile, req)

		var filter newModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListNew(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
