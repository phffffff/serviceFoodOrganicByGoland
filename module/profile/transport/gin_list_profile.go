package profileTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	profileBusiness "go_service_food_organic/module/profile/business"
	profileModel "go_service_food_organic/module/profile/model"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinListProfile(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()

		req := c.MustGet(common.CurrentUser).(common.Requester)

		var filter profileModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		store := profileStorage.NewSqlModel(db)
		biz := profileBusiness.NewListProfileBiz(store, req)

		list, err := biz.ListProfileWithFilter(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mark(false)
			if list[i].Image != nil {
				list[i].Image.Mark(false)
			}
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
