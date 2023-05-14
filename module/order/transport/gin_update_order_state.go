package orderTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodRepo "go_service_food_organic/module/food/repository"
	foodStorage "go_service_food_organic/module/food/storage"
	orderBusiness "go_service_food_organic/module/order/business"
	orderModel "go_service_food_organic/module/order/model"
	orderRepo "go_service_food_organic/module/order/repository"
	orderStorage "go_service_food_organic/module/order/storage"
	"net/http"
)

func GinUpdateOrderState(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()

		store := orderStorage.NewSqlModel(db)
		storeFood := foodStorage.NewSqlModel(db)

		req := c.MustGet(common.CurrentUser).(common.Requester)

		foodUpdateCountRepo := foodRepo.NewUpdateFoodCountRepo(storeFood)
		repo := orderRepo.NewUpdateOrderStateRepo(store, foodUpdateCountRepo, req)

		biz := orderBusiness.NewUpdateOrderStateBiz(repo)

		var orderUpdate orderModel.OrderUpdate
		if err := c.ShouldBind(&orderUpdate); err != nil {
			panic(err)
		}

		if err := biz.UpdateOrderState(c.Request.Context(), id, orderUpdate.State); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
