package appContext

import "gorm.io/gorm"

type AppContext interface {
	GetMyDBConnection() *gorm.DB
}

type appContext struct {
	db *gorm.DB
}

func NewAppContext(db *gorm.DB) *appContext {
	return &appContext{db: db}
}

func (appCtx *appContext) GetMyDBConnection() *gorm.DB {
	return appCtx.db
}
