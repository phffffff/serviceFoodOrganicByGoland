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

func GinCreateBrand(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := brandStorage.NewSqlModel(db)
		biz := brandBusiness.NewCreateBrandBiz(store, req)

		var data brandModel.BrandCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.CreateBrand(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
