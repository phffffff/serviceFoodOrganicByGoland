package userTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	hash "go_service_food_organic/component/hasher"
	userBusiness "go_service_food_organic/module/user/business"
	userModel "go_service_food_organic/module/user/model"
	userRepo "go_service_food_organic/module/user/repository"
	userStorage "go_service_food_organic/module/user/storage"
	"net/http"
)

func GinUpdateUser(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		var data userModel.UserPasswordUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}
		req, ok := c.MustGet(common.CurrentUser).(common.Requester)
		if !ok {
			panic(common.ErrorNoPermission(nil))
		}

		db := appCtx.GetMyDBConnection()
		hasher := hash.NewMd5Hash()
		store := userStorage.NewSqlModel(db)
		repo := userRepo.NewUpdatePassRepo(store, hasher, req)
		biz := userBusiness.NewUpdatePassBiz(repo)
		if err := biz.UpdateUserPassword(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))

	}
}
