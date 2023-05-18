package aboutTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	aboutBusiness "go_service_food_organic/module/about/business"
	aboutModel "go_service_food_organic/module/about/model"
	aboutRepo "go_service_food_organic/module/about/repository"
	aboutStorage "go_service_food_organic/module/about/storage"
	"net/http"
)

func GinUpdateAbout(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := aboutStorage.NewSqlModel(db)
		repo := aboutRepo.NewUpdateAboutRepo(store, req)
		biz := aboutBusiness.NewUpdateAboutBiz(repo)

		var data aboutModel.AboutUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.UpdateAbout(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
