package newTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	newBusiness "go_service_food_organic/module/new/business"
	newModel "go_service_food_organic/module/new/model"
	newRepo "go_service_food_organic/module/new/repository"
	newStorage "go_service_food_organic/module/new/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinUpdateNew(appCtx appContext.AppContext) gin.HandlerFunc {
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
		repo := newRepo.NewUpdateNewRepo(store, storeProfile, req)
		biz := newBusiness.NewUpdateNewBiz(repo)

		var data newModel.NewUpd

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.UpdateNew(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
