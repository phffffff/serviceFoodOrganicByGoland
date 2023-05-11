package userTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	userBusiness "go_service_food_organic/module/user/business"
	userModel "go_service_food_organic/module/user/model"
	userStorage "go_service_food_organic/module/user/storage"
	"net/http"
)

func GinListUser(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()

		req := c.MustGet(common.CurrentUser).(common.Requester)

		var filter userModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		filter.Status = []int{1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		store := userStorage.NewSqlModel(db)
		biz := userBusiness.NewListUserBiz(store, req)

		list, err := biz.ListUserWithFilter(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mark(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
