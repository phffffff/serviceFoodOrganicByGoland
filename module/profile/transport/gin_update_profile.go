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
	"strconv"
)

func GinUpdateProfile(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			panic(err)
		}

		var data profileModel.Profile

		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		store := profileStorage.NewSqlModel(db)
		repo := profileRepo.NewUpdateProfileRepo(store)
		biz := profileBusiness.NewUpdateProfileBiz(repo)

		if err := biz.UpdateProfile(c.Request.Context(), id, &data); err != nil {
			panic(err)
		}
		data.Mark(false)
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(data.FakeId.String()))

	}
}
