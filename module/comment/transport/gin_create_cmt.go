package commentTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	commentBusiness "go_service_food_organic/module/comment/business"
	commentModel "go_service_food_organic/module/comment/model"
	commentStorage "go_service_food_organic/module/comment/storage"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinCreateCmt(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		newUID, err := common.FromBase58(c.Param("new_id"))
		if err != nil {
			panic(err)
		}
		newId := int(newUID.GetLocalID())

		db := appCtx.GetMyDBConnection()
		req := c.MustGet(common.CurrentUser).(common.Requester)
		storeProfile := profileStorage.NewSqlModel(db)
		store := commentStorage.NewSqlModel(db)

		biz := commentBusiness.NewCreateCmtBiz(store, storeProfile, req)

		var data commentModel.CommentCrt
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}

		data.NewId = newId

		if err := biz.CreateCmt(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
