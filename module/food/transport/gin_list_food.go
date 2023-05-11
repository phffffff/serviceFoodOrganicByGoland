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

func GinListFood(appctx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appctx.GetMyDBConnection()

		store := foodStorage.NewSqlModel(db)
		biz := foodBusiness.NewListFoodBiz(store)

		var filter foodModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		filter.Status = []int{0, 1}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		paging.FullFill()

		list, err := biz.ListFoodWithFilter(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}
		for i := range list {
			list[i].Mark(false)
			for _, item := range list[i].FoodImages {
				item.Mark(false)
			}
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))

	}
}
