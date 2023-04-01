package foodTransport

import (
	"github.com/gin-gonic/gin"
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
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		db := appCtx.GetMyDBConnection()
		store := foodStorage.NewSqlModel(db)
		biz := foodBusiness.NewCreateFoodBiz(store)

		if err := biz.CreateFood(c.Request.Context(), &data); err != nil {
			c.IndentedJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"data": data.Id,
		})

	}
}
