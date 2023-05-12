package orderDetailTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodStorage "go_service_food_organic/module/food/storage"
	orderDetailBusiness "go_service_food_organic/module/order_detail/business"
	orderDetailModel "go_service_food_organic/module/order_detail/model"
	orderDetailRepo "go_service_food_organic/module/order_detail/repository"
	orderDetailStorage "go_service_food_organic/module/order_detail/storage"
	"net/http"
)

func GinCreateOrderDetail(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := orderDetailStorage.NewSqlModel(db)
		storeFood := foodStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := orderDetailRepo.NewCreateOrderDetailRepo(store, storeFood, req)
		biz := orderDetailBusiness.NewCreateOrderDetailBiz(repo)

		var od orderDetailModel.OrderDetailCreate
		if err := c.ShouldBind(&od); err != nil {
			panic(err)
		}

		foodUID, err := common.FromBase58(od.FoodFakeId)
		if err != nil {
			panic(err)
		}

		orderUID, err := common.FromBase58(od.OrderFakeId)
		if err != nil {
			panic(err)
		}

		od.FoodId = int(foodUID.GetLocalID())
		od.OrderId = int(orderUID.GetLocalID())

		if err := biz.CreateOrderDetail(c.Request.Context(), &od); err != nil {
			panic(err)
		}

		od.Mark(false)
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(od.FakeId.String()))
	}
}
