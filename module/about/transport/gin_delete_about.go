package aboutTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	aboutBusiness "go_service_food_organic/module/about/business"
	aboutRepo "go_service_food_organic/module/about/repository"
	aboutStorage "go_service_food_organic/module/about/storage"
	"net/http"
)

func GinDeleteAbout(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := aboutStorage.NewSqlModel(db)
		repo := aboutRepo.NewDeleteAboutRepo(store, req)
		biz := aboutBusiness.NewDeleteAboutBiz(repo)

		if err := biz.DeleteAbout(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
