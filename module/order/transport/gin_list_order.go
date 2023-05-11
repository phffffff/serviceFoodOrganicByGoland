package orderTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	orderBusiness "go_service_food_organic/module/order/business"
	orderModel "go_service_food_organic/module/order/model"
	orderRepo "go_service_food_organic/module/order/repository"
	orderStorage "go_service_food_organic/module/order/storage"
	"net/http"
)

func GinListOrder(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := orderStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := orderRepo.NewListOrderRepo(store, req)
		biz := orderBusiness.NewListOrderBiz(repo)

		var filter orderModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		filter.Status = []int{1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListOrder(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mark(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
