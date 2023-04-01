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
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		paging.FullFill()

		list, err := biz.ListFoodWithFilter(c.Request.Context(), &filter, &paging)
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusOK, gin.H{
			"data": list,
		})

	}
}
