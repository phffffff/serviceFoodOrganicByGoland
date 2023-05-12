package cartTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	cartBusiness "go_service_food_organic/module/carts/business"
	cartRepo "go_service_food_organic/module/carts/repository"
	cartStorage "go_service_food_organic/module/carts/storage"
	"net/http"
)

func GinDeleteCart(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := cartStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := cartRepo.NewDeleteCartRepo(store, req)
		biz := cartBusiness.NewDeleteFoodBiz(repo)

		if err := biz.DeleteFood(c.Request.Context()); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
