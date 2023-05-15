package infoFoodcategoryTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	infoFoodCategoryBusiness "go_service_food_organic/module/info_food_category/business"
	infoFoodCategoryModel "go_service_food_organic/module/info_food_category/model"
	infoFoodCategoryRepo "go_service_food_organic/module/info_food_category/repository"
	infoFoodCategoryStorage "go_service_food_organic/module/info_food_category/storage"
	"net/http"
)

func GinUpdateInfoFoodCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()
		store := infoFoodCategoryStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)

		repo := infoFoodCategoryRepo.NewUpdateInfoFoodCategoryRepo(store, req)
		biz := infoFoodCategoryBusiness.NewUpdateInfoFoodCategoryBiz(repo)

		var data infoFoodCategoryModel.InfoFoodCategoryUpdate
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

		data.FoodFakeId = ""
		data.CategoryFakeId = ""

		if err := biz.UpdateInfoFoodCategory(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
