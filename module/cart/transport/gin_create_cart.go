package cartTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	cartBusiness "go_service_food_organic/module/cart/business"
	cartModel "go_service_food_organic/module/cart/model"
	cartRepo "go_service_food_organic/module/cart/repository"
	cartStorage "go_service_food_organic/module/cart/storage"
	"net/http"
)

func GinCreateCart(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := cartStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := cartRepo.NewCreateCartRepo(store, req)
		biz := cartBusiness.NewCreateCartBiz(repo)

		var cart cartModel.Cart
		if err := c.ShouldBind(&cart); err != nil {
			panic(err)
		}

		uid, err := common.FromBase58(cart.FoodFakeId)
		if err != nil {
			panic(err)
		}
		cart.FoodId = int(uid.GetLocalID())

		if err := biz.CreateCart(c.Request.Context(), &cart); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
