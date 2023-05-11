package orderDetailTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	orderDetailBusiness "go_service_food_organic/module/order_detail/business"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
	orderDetailRepo "go_service_food_organic/module/order_detail/repository"
	orderDetailStorage "go_service_food_organic/module/order_detail/storage"
	"net/http"
)

func GinListOrderDetail(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := orderDetailStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := orderDetailRepo.NewListOrderDetailRepo(store, req)
		biz := orderDetailBusiness.NewListOrderDetailBiz(repo)

		var filter orderDetailModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}
		filter.Status = []int{1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListOrderDetail(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mark(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
