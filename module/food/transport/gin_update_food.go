package foodTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodBusiness "go_service_food_organic/module/food/business"
	foodModel "go_service_food_organic/module/food/model"
	foodRepo "go_service_food_organic/module/food/repository"
	foodStorage "go_service_food_organic/module/food/storage"
	"net/http"
)

func GinUpdateFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			panic(err)
		}
		id := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()
		store := foodStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := foodRepo.NewUpdateFoodRepo(store, req)
		biz := foodBusiness.NewUpdateFoodBiz(repo)

		var data foodModel.FoodUpdate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		brandUid, err := common.FromBase58(data.BrandFakeId)
		if err != nil {
			panic(err)
		}

		data.BrandId = int(brandUid.GetLocalID())

		if err := biz.UpdateFood(c.Request.Context(), &data, id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
