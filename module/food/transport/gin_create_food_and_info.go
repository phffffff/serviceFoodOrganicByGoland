package foodTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodBusiness "go_service_food_organic/module/food/business"
	foodModel "go_service_food_organic/module/food/model"
	foodRepo "go_service_food_organic/module/food/repository"
	foodStorage "go_service_food_organic/module/food/storage"
	infoFoodCategoryStorage "go_service_food_organic/module/info_food_category/storage"
	"net/http"
)

func GinCreateFoodAndInfo(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("categoryId"))
		if err != nil {
			panic(err)
		}
		categoryId := int(uid.GetLocalID())

		db := appCtx.GetMyDBConnection()
		storeFood := foodStorage.NewSqlModel(db)
		storeInfo := infoFoodCategoryStorage.NewSqlModel(db)
		req := c.MustGet(common.CurrentUser).(common.Requester)
		repo := foodRepo.NewCreateFoodAndInfo(storeFood, storeInfo, req)
		biz := foodBusiness.NewCreateFoodAndInfoBiz(repo)

		var data foodModel.FoodCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		if err := biz.CreateFoodAndInfo(c.Request.Context(), &data, categoryId); err != nil {
			panic(err)
		}
		data.Mark(false)
		c.IndentedJSON(http.StatusBadRequest, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
