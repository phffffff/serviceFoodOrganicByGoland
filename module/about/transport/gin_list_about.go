package aboutTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	aboutBusiness "go_service_food_organic/module/about/business"
	aboutModel "go_service_food_organic/module/about/model"
	aboutStorage "go_service_food_organic/module/about/storage"
	"net/http"
)

func GinListAbout(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := aboutStorage.NewSqlModel(db)
		biz := aboutBusiness.NewlistAboutBiz(store)

		var filter aboutModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListAbout(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
