package commentTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	commentBusiness "go_service_food_organic/module/comment/business"
	commentModel "go_service_food_organic/module/comment/model"
	commentStorage "go_service_food_organic/module/comment/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinListCmt(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newUID, err := common.FromBase58(c.Param("new_id"))
		if err != nil {
			panic(err)
		}
		newId := int(newUID.GetLocalID())

		db := appCtx.GetMyDBConnection()
		store := commentStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		biz := commentBusiness.NewListCmtBiz(store, storeProfile, req)

		var filter commentModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		filter.NewId = newId

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListCmt(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
