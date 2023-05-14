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

func GinListCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := categoryStorage.NewSqlModel(db)
		biz := categoryBusiness.NewlistCategoryBiz(store)

		var filter categoryModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListCategory(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
			for _, item := range list[i].Foods {
				item.Mark(false)
			}
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
