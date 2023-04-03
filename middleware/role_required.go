package middleware

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
)

func RoleRequired(appCtx appContext.AppContext, allowRole ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet(common.CurrentUser).(common.Requester)

		var hasFound bool = false
		for idx := range allowRole {
			if allowRole[idx] == user.GetRole() {
				hasFound = true
				break
			}
		}
		if !hasFound {
			panic(common.ErrorNoPermission(nil))
		}

		c.Next()
	}
}
