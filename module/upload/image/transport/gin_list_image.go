package imageTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	imageBusiness "go_service_food_organic/module/upload/image/business"
	imageModel "go_service_food_organic/module/upload/image/model"
	imageStorage "go_service_food_organic/module/upload/image/storage"
	"net/http"
)

func GinListImage(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		var filter imageModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		filter.Status = []int{0, 1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := imageStorage.NewSqlModel(db)
		biz := imageBusiness.NewListImageBiz(store, req)

		list, err := biz.ListImage(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		for idx := range list {
			list[idx].Mark(false)
		}
		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
