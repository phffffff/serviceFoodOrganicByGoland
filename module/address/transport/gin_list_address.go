package addressTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	addressBusiness "go_service_food_organic/module/address/business"
	addressModel "go_service_food_organic/module/address/model"
	addressStorage "go_service_food_organic/module/address/storage"
	"net/http"
)

func GinListAddress(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := addressStorage.NewSqlModel(db)
		biz := addressBusiness.NewlistAddressBiz(store)

		var filter addressModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListAddress(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
			list[i].Provinces.Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
