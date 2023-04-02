package userTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/component/hasher"
	userBusiness "go_service_food_organic/module/user/business"
	userModel "go_service_food_organic/module/user/model"
	userStorage "go_service_food_organic/module/user/storage"
	"net/http"
)

func GinRegister(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userModel.UserRegister

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		hasher := hash.NewMd5Hash()
		store := userStorage.NewSqlModel(db)
		biz := userBusiness.NewRegisterBiz(store, hasher)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.Id))

	}
}
