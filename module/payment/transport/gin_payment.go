package paymentTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	cartRepo "go_service_food_organic/module/carts/repository"
	cartStorage "go_service_food_organic/module/carts/storage"
	foodRepo "go_service_food_organic/module/food/repository"
	foodStorage "go_service_food_organic/module/food/storage"
	orderRepo "go_service_food_organic/module/order/repository"
	orderStorage "go_service_food_organic/module/order/storage"
	orderDetailStorage "go_service_food_organic/module/order_detail/storage"
	paymentBusiness "go_service_food_organic/module/payment/business"
	paymentRepo "go_service_food_organic/module/payment/repository"
	"net/http"
)

func GinPayment(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()

		req := c.MustGet(common.CurrentUser).(common.Requester)

		storeCart := cartStorage.NewSqlModel(db)
		storeFood := foodStorage.NewSqlModel(db)
		storeOrder := orderStorage.NewSqlModel(db)
		storeOrderUpdate := orderStorage.NewSqlModel(db)
		storeFoodUpdate := foodStorage.NewSqlModel(db)
		storeOrderDetail := orderDetailStorage.NewSqlModel(db)
		updateOrderPriceRepo := orderRepo.NewUpdateOrderPriceRepo(storeOrderUpdate)
		updateFoodCountRepo := foodRepo.NewUpdateFoodCountRepo(storeFoodUpdate)
		deleteCartWhenRepo := cartRepo.NewDeleteCartRepo(storeCart, req)

		repo := paymentRepo.NewPaymentRepo(
			storeCart,
			storeFood,
			storeOrder,
			storeOrderDetail,
			updateOrderPriceRepo,
			updateFoodCountRepo,
			deleteCartWhenRepo,
			req)
		biz := paymentBusiness.NewPaymentBiz(repo)

		if err := biz.Payment(c); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse("success"))
	}
}
