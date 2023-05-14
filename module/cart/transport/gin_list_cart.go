package cartTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	cartBusiness "go_service_food_organic/module/cart/business"
	cartModel "go_service_food_organic/module/cart/model"
	cartStorage "go_service_food_organic/module/cart/storage"
	"net/http"
)

func GinListCart(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMyDBConnection()
		store := cartStorage.NewSqlModel(db)
		biz := cartBusiness.NewListCartBiz(store)

		req := c.MustGet(common.CurrentUser).(common.Requester)

		var filter cartModel.Filter
		if err := c.ShouldBind(&filter); err != nil {
			panic(err)
		}

		filter.UserId = req.GetUserId()

		var paging common.Paging
		if err := c.ShouldBind(&paging); err != nil {
			panic(err)
		}
		paging.FullFill()

		list, err := biz.ListCart(c.Request.Context(), &filter, &paging)
		if err != nil {
			panic(err)
		}

		for i := range list {
			list[i].Mask(false)
		}

		c.IndentedJSON(http.StatusOK, common.FullSuccessResponse(list, filter, paging))
	}
}
