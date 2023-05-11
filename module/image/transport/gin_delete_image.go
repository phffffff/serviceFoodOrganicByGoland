package imageTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/module/image/business"
	"go_service_food_organic/module/image/repository"
	"go_service_food_organic/module/image/storage"
	"net/http"
)

func GinDeleteImage(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()

		req := c.MustGet(common.CurrentUser).(common.Requester)

		store := imageStorage.NewSqlModel(db)
		repo := imageRepo.NewDeleteImageRepo(store, req)
		biz := imageBusiness.NewDeleteImageBiz(repo)

		if err := biz.DeteleImage(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
