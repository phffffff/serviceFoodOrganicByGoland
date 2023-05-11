package foodTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodBusiness "go_service_food_organic/module/food/business"
	foodRepo "go_service_food_organic/module/food/repository"
	foodStorage "go_service_food_organic/module/food/storage"
	"net/http"
)

func GinDeleteFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()

		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := foodStorage.NewSqlModel(db)
		repo := foodRepo.NewDeleteFoodRepo(store, req)
		biz := foodBusiness.NewDeleteFoodBiz(repo)

		if err := biz.DeleteFood(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
