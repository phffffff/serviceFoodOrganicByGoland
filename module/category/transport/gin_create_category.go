package categoryTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	categoryBusiness "go_service_food_organic/module/category/business"
	categoryModel "go_service_food_organic/module/category/model"
	categoryStorage "go_service_food_organic/module/category/storage"
	"net/http"
)

func GinCreateCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := categoryStorage.NewSqlModel(db)
		biz := categoryBusiness.NewCreateCategoryBiz(store, req)

		var data categoryModel.CategoryCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.CreateCategory(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
