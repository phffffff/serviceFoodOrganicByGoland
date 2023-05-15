package categoryTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	categoryBusiness "go_service_food_organic/module/category/business"
	categoryRepo "go_service_food_organic/module/category/repository"
	categoryStorage "go_service_food_organic/module/category/storage"
	"net/http"
)

func GinDeleteCategory(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := common.FromBase58(c.Param("id"))
		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, common.ErrInvalidRequest(err))
		}
		id := int(uid.GetLocalID())

		req := c.MustGet(common.CurrentUser).(common.Requester)

		db := appCtx.GetMyDBConnection()
		store := categoryStorage.NewSqlModel(db)
		repo := categoryRepo.NewDeleteCategoryRepo(store, req)
		biz := categoryBusiness.NewDeleteCategoryBiz(repo)

		if err := biz.DeleteCategory(c.Request.Context(), id); err != nil {
			panic(err)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(true))
	}
}
