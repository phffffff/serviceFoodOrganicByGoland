package appContext

import (
	uploadProvider "go_service_food_organic/component/upload_provider"
	"gorm.io/gorm"
)

type AppContext interface {
	GetSecretSaltHashImage() string
	GetMyDBConnection() *gorm.DB
	GetSecretkey() string
	UploadProvider() uploadProvider.UploadProvider
}

type appContext struct {
	db             *gorm.DB
	secretKey      string
	uploadProvider uploadProvider.UploadProvider
	secretSalt     string
}

func NewAppContext(
	db *gorm.DB,
	secretKey string,
	uploadProvider uploadProvider.UploadProvider,
	secretSalt string,
) *appContext {
	return &appContext{
		db:             db,
		secretKey:      secretKey,
		uploadProvider: uploadProvider,
		secretSalt:     secretKey,
	}
}

func (appCtx *appContext) GetMyDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) GetSecretkey() string {
	return appCtx.secretKey
}

func (appCtx *appContext) UploadProvider() uploadProvider.UploadProvider {
	return appCtx.uploadProvider
}

func (appCtx *appContext) GetSecretSaltHashImage() string {
	return appCtx.secretSalt
}
