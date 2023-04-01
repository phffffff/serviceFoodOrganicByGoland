package main

import (
	"github.com/gin-gonic/gin"
	appContext "go_service_food_organic/component/app_context"
	foodTransport "go_service_food_organic/module/food/transport"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "cool_organic:@Klov3x124n@tcp(127.0.0.1:3307)/cool_organic?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err)
	}

	db.Debug()

	appCtx := appContext.NewAppContext(db)

	rt := gin.Default()
	{
		food := rt.Group("food")
		food.GET("/listfood", foodTransport.GinListFood(appCtx))

		food.POST("/createfood", foodTransport.GinCreateFood(appCtx))
	}

	rt.Run()
}
