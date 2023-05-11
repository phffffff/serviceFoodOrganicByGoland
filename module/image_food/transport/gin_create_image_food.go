package imageFoodTransport

import (
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	foodStorage "go_service_food_organic/module/food/storage"
	imageStorage "go_service_food_organic/module/image/storage"
	"go_service_food_organic/module/image_food/business"
	imageFoodModel "go_service_food_organic/module/image_food/model"
	imageFoodRepo "go_service_food_organic/module/image_food/repository"
	"go_service_food_organic/module/image_food/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GinCreateImageFood(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := c.MustGet(common.CurrentUser).(common.Requester)
		db := appCtx.GetMyDBConnection()
		store := imageFoodStorage.NewSqlModel(db)
		storeImage := imageStorage.NewSqlModel(db)
		storeFood := foodStorage.NewSqlModel(db)
		repo := imageFoodRepo.NewCreateImageFoodRepo(store, req, storeImage, storeFood)
		biz := imageFoodBusiness.NewCreateImageFoodBiz(repo)

		var imf imageFoodModel.ImageFoodCreate
		if err := c.ShouldBind(&imf); err != nil {
			panic(err)
		}

		foodFakeId, err := common.FromBase58(imf.FoodFakeId)
		if err != nil {
			panic(err)
		}
		imageFakeId, err := common.FromBase58(imf.ImageFakeId)
		if err != nil {
			panic(err)
		}

		imf.FoodId = int(foodFakeId.GetLocalID())
		imf.ImageId = int(imageFakeId.GetLocalID())

		if err := biz.CreateImageFood(c.Request.Context(), &imf); err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
