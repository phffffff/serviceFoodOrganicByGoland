package infoFoodcategoryTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	infoFoodCategoryBusiness "go_service_food_organic/module/info_food_category/business"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
	infoFoodCategoryStorage "go_service_food_organic/module/info_food_category/storage"
	"net/http"
)

func GinCreateInfoFoodCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMyDBConnection()
		store := infoFoodCategoryStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)

		biz := infoFoodCategoryBusiness.NewCreateInfoFoodCategoryBiz(store, req)

		var data infoFoodCategoryModel.InfoFoodCategoryCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		foodUID, err := common.FromBase58(data.FoodFakeId)
		if err != nil {
			panic(err)
		}

		categoryUID, err := common.FromBase58(data.CategoryFakeId)
		if err != nil {
			panic(err)
		}

		data.FoodId = int(foodUID.GetLocalID())
		data.CategoryId = int(categoryUID.GetLocalID())

		if err := biz.CreateInfoFoodCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
