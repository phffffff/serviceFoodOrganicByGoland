package foodTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodBusiness "go_service_food_organic/module/food/business"
	foodModel "go_service_food_organic/module/food/model"
	foodStorage "go_service_food_organic/module/food/storage"
	"net/http"
)

func GinCreateFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data foodModel.FoodCreate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		store := foodStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		biz := foodBusiness.NewCreateFoodBiz(store, req)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mark(false)
		c.IndentedJSON(http.StatusBadRequest, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
