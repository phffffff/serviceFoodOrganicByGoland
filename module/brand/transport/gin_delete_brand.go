package brandTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	brandBusiness "go_service_food_organic/module/brand/business"
	brandRepo "go_service_food_organic/module/brand/repository"
	brandStorage "go_service_food_organic/module/brand/storage"
	"net/http"
)

func GinDeleteBrand(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := brandStorage.NewSqlModel(db)
		repo := brandRepo.NewDeleteBrandRepo(store, req)
		biz := brandBusiness.NewDeleteBrandBiz(repo)

		if err := biz.DeleteBrand(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
