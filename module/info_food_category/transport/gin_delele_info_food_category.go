package infoFoodcategoryTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	infoFoodCategoryBusiness "go_service_food_organic/module/info_food_category/business"
	infoFoodCategoryRepo "go_service_food_organic/module/info_food_category/repository"
	infoFoodCategoryStorage "go_service_food_organic/module/info_food_category/storage"
	"net/http"
)

func GinDeleteInfoFoodCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()
		store := infoFoodCategoryStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)

		repo := infoFoodCategoryRepo.NewDeleteInfoFoodCategoryRepo(store, req)
		biz := infoFoodCategoryBusiness.NewDeleteInfoFoodCategoryBiz(repo)

		if err := biz.DeleteInfoFoodCategory(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
