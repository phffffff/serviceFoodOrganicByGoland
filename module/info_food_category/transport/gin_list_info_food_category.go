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

func GinListInfoFoodCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		db := appCtx.GetMyDBConnection()
		store := infoFoodCategoryStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)

		biz := infoFoodCategoryBusiness.NewListInfoFoodCategoryBiz(store, req)

		var filter infoFoodCategoryModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListInfoFoodCategory(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
			list[i].Foods.Mark(false)
			list[i].Categories.Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
