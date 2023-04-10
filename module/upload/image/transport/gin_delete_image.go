package imageTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	imageBusiness "go_service_food_organic/module/upload/image/business"
	imageRepo "go_service_food_organic/module/upload/image/repository"
	imageStorage "go_service_food_organic/module/upload/image/storage"
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

		store := imageStorage.NewSqlModel(db)
		repo := imageRepo.NewDeleteImageRepo(store)
		biz := imageBusiness.NewDeleteImageBiz(repo)

		if err := biz.DeteleImage(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
