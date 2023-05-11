package imageFoodTransport

import (
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/module/image_food/business"
	"go_service_food_organic/module/image_food/repository"
	"go_service_food_organic/module/image_food/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinDeleteImageFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}

		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := imageFoodStorage.NewSqlModel(db)
		repo := imageFoodRepo.NewDeleteImageFoodRepo(store, req)
		biz := imageFoodBusiness.NewDeleteImageFoodBiz(repo)

		if err := biz.DeleteImageFood(c.Request.Context(), id); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
