package imageFoodTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	imageFoodBusiness "go_service_food_organic/module/image_food/business"
	imageFoodModel "go_service_food_organic/module/image_food/model"
	imageFoodStorage "go_service_food_organic/module/image_food/storage"
	"net/http"
)

func GinListImageFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		req := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMyDBConnection()
		store := imageFoodStorage.NewSqlModel(db)
		biz := imageFoodBusiness.NewListImageFoodBiz(store, req)

		var filter imageFoodModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		//filter.Status = []int{0, 1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListImageFood(c.Request.Context(), &filter, &paging)
		if err != nil && list == nil {
			panic(err)
		}

		for idx := range list {
			list[idx].Mark(false)
		}
		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
