package userTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	hash "go_service_food_organic/component/hasher"
	"go_service_food_organic/component/token/jwt"
	userBusiness "go_service_food_organic/module/user/business"
	userModel "go_service_food_organic/module/user/model"
	userStorage "go_service_food_organic/module/user/storage"
	"net/http"
)

func GinLogin(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data userModel.UserLogin

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		db := appCtx.GetMyDBConnection()
		secretSalt := appCtx.GetSecretSaltHashImage()

		store := userStorage.NewSqlModel(db)
		hasher := hash.NewMd5Hash(secretSalt)
		tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretkey())
		biz := userBusiness.NewLoginBiz(store, hasher, tokenProvider, 60*60*24*30)

		token, err := biz.Login(c.Request.Context(), &data)
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(token))
	}
}
