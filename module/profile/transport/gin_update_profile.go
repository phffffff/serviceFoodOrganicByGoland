package profileTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	profileBusiness "go_service_food_organic/module/profile/business"
	profileModel "go_service_food_organic/module/profile/model"
	profileRepo "go_service_food_organic/module/profile/repository"
	profileStorage "go_service_food_organic/module/profile/storage"
	"net/http"
)

func GinUpdateProfile(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {

		req := c.MustGet(common.CurrentUser).(common.Requester)

		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		var data profileModel.ProfileUpdate

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		store := profileStorage.NewSqlModel(db)
		repo := profileRepo.NewUpdateProfileRepo(store, req)
		biz := profileBusiness.NewUpdateProfileBiz(repo)

		if err := biz.UpdateProfile(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
