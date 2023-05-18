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

func GinCreateAbout(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		store := aboutStorage.NewSqlModel(db)
		biz := aboutBusiness.NewCreateAboutBiz(store, req)

		var data aboutModel.About
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.CreateAbout(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
