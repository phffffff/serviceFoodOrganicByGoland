package userTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/component/hasher"
	profileStorage "go_service_food_organic/module/profile/storage"
	userBusiness "go_service_food_organic/module/user/business"
	userModel "go_service_food_organic/module/user/model"
	userRepo "go_service_food_organic/module/user/repository"
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
		secretSalt := appCtx.GetSecretSaltHashImage()

		hasher := hash.NewMd5Hash(secretSalt)
		storeUser := userStorage.NewSqlModel(db)
		storeProfile := profileStorage.NewSqlModel(db)
		repo := userRepo.NewRegisterRepo(storeUser, hasher, storeProfile)
		biz := userBusiness.NewRegisterBiz(repo)

		if err := biz.Register(c.Request.Context(), &data); err != nil {
			panic(err)
		}

		data.Mark(false)

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
