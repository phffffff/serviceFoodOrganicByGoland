package provinceTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	provinceBusiness "go_service_food_organic/module/province/business"
	provinceModel "go_service_food_organic/module/province/model"
	provinceStorage "go_service_food_organic/module/province/storage"
	"net/http"
)

func GinListProvince(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := provinceStorage.NewSqlModel(db)
		biz := provinceBusiness.NewListProvinceBiz(store)

		var filter provinceModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}

		list, err := biz.ListProvince(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, panic))
	}
}
