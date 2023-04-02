package appContext

import "gorm.io/gorm"

type AppContext interface {
	GetMyDBConnection() *gorm.DB
	GetSecretkey() string
}

type appContext struct {
	db        *gorm.DB
	secretKey string
}

func NewAppContext(db *gorm.DB, secretKey string) *appContext {
	return &appContext{
		db:        db,
		secretKey: secretKey,
	}
}

func (appCtx *appContext) GetMyDBConnection() *gorm.DB {
	return appCtx.db
}

func (appCtx *appContext) GetSecretkey() string {
	return appCtx.secretKey
}
