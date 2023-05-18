package commentTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	commentBusiness "go_service_food_organic/module/comment/business"
	commentModel "go_service_food_organic/module/comment/model"
	commentRepo "go_service_food_organic/module/comment/repository"
	commentStorage "go_service_food_organic/module/comment/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinUpdateCmt(appCtx appContext.AppContext) gin.HandlerFunc {
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
		repo := commentRepo.NewUpdateCmtRepo(store, storeProfile, req)
		biz := commentBusiness.NewUpdateCmtBiz(repo)

		var data commentModel.CommentUpd

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.UpdateCmt(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
