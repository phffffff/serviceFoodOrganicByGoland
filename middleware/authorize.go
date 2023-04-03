package middleware

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	"go_service_food_organic/component/token/jwt"
	userModel "go_service_food_organic/module/user/model"
	userStorage "go_service_food_organic/module/user/storage"
	"strings"
)

const (
	ErrWrongAuthHeader = "ErrWrongAuthHeader"
	MsgWrongAuthHeader = "wrong authen header"
)

func ErrorWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(err, MsgWrongAuthHeader, ErrWrongAuthHeader)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")

	if parts[0] != "Bearer" || len(parts) > 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrorWrongAuthHeader(nil)
	}
	return parts[1], nil
}

func RequiredAuth(appCtx appContext.AppContext) gin.HandlerFunc {
	tokenProvider := jwt.NewTokenJWTProvider(appCtx.GetSecretkey())

	return func(c *gin.Context) {
		token, err := extractTokenFromHeaderString(c.GetHeader("Authorization"))

		if err != nil {
			panic(err)
		}

		db := appCtx.GetMyDBConnection()
		store := userStorage.NewSqlModel(db)

		payload, err := tokenProvider.Validate(token)
		if err != nil {
			panic(err)
		}
		user, err := store.FindDataWithCondition(
			c.Request.Context(),
			map[string]interface{}{"id": payload.UserId},
		)

		if err != nil {
			panic(err)
		}

		if user.Status == 0 {
			panic(userModel.ErrorUserExists())
		}

		user.Mark(false)

		c.Set(common.CurrentUser, user)
		c.Next()
	}
}
