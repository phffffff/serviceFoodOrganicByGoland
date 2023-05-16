package brandTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	brandBusiness "go_service_food_organic/module/brand/business"
	brandModel "go_service_food_organic/module/brand/model"
	brandStorage "go_service_food_organic/module/brand/storage"
	"net/http"
)

func GinListBrand(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := brandStorage.NewSqlModel(db)
		biz := brandBusiness.NewlistBrandBiz(store)

		var filter brandModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListBrand(c.Request.Context(), &filter, &paging)
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
