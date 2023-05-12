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

func GinCreateOrder(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := orderStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := orderRepo.NewCreateOrderRepo(store, req)
		biz := orderBusiness.NewCreateOrderBiz(repo)

		var order orderModel.OrderCreate

		if err := c.ShouldBind(&order); err != nil {
			panic(err)
		}

		userUID, err := common.FromBase58(order.UserFakeId)
		if err != nil {
			panic(err)
		}
		order.UserId = int(userUID.GetLocalID())

		if err := biz.CreateOrder(c.Request.Context(), &order); err != nil {
			panic(err)
		}

		order.Mark(false)
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(order.FakeId.String()))
	}
}
