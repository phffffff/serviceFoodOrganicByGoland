package newTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	newBusiness "go_service_food_organic/module/new/business"
	newRepo "go_service_food_organic/module/new/repository"
	newStorage "go_service_food_organic/module/new/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinDeleteNew(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := newStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		repo := newRepo.NewDeleteNewRepo(store, storeProfile, req)
		biz := newBusiness.NewDeleteNewBiz(repo)

		if err := biz.DeleteNew(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
