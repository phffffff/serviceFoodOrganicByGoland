package profileStorage

import "gorm.io/gorm"

type sqlModel struct {
	db *gorm.DB
}

func NewSqlModel(db *gorm.DB) *sqlModel {
	return &sqlModel{db: db}
}
