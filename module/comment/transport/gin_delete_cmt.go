package commentTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	commentBusiness "go_service_food_organic/module/comment/business"
	commentRepo "go_service_food_organic/module/comment/repository"
	commentStorage "go_service_food_organic/module/comment/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinDeleteCmt(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := commentStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		repo := commentRepo.NewDeleteCmtRepo(store, storeProfile, req)
		biz := commentBusiness.NewDeleteCmtBiz(repo)

		if err := biz.DeleteCmt(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
