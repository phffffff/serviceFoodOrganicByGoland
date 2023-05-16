package main

import (
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	uploadProvider "go_service_food_organic/component/upload_provider"
	"go_service_food_organic/middleware"
	brandTransport "go_service_food_organic/module/brand/transport"
	cartTransport "go_service_food_organic/module/cart/transport"
	categoryTransport "go_service_food_organic/module/category/transport"
	foodTransport "go_service_food_organic/module/food/transport"
	imageTransport "go_service_food_organic/module/image/transport"
	imageFoodTransport "go_service_food_organic/module/image_food/transport"
	infoFoodcategoryTransport "go_service_food_organic/module/info_food_category/transport"
	orderTransport "go_service_food_organic/module/order/transport"
	orderDetailTransport "go_service_food_organic/module/order_detail/transport"
	paymentTransport "go_service_food_organic/module/payment/transport"
	profileTransport "go_service_food_organic/module/profile/transport"
	provinceTransport "go_service_food_organic/module/province/transport"
	userTransport "go_service_food_organic/module/user/transport"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := os.Getenv("DB_CONN")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("err:", err)
	}

	db.Debug()

	secretKey := os.Getenv("SYSTEM_SECRET")
	secretSalt := os.Getenv("SALT_HASH_DATA_IMG")

	s3BucketName := os.Getenv("S3_BUCKET_NAME")
	s3Region := os.Getenv("S3_REGION")
	s3APIKey := os.Getenv("S3_ACCESS_KEY")
	s3SecretKey := os.Getenv("S3_SECRET_KEY")
	s3Domain := os.Getenv("S3_DOMAIN")

	s3Provider := uploadProvider.NewS3Provider(s3BucketName, s3Region, s3APIKey, s3SecretKey, s3Domain)

	appCtx := appContext.NewAppContext(db, secretKey, s3Provider, secretSalt)

	rt := gin.Default()
	rt.Use(middleware.Recover(appCtx))

	{
		admin := rt.Group(
			"/admin",
			middleware.RequiredAuth(appCtx),
			middleware.RoleRequired(appCtx, common.Admin),
		)

		{
			upload := admin.Group("image")
			upload.POST("upload", imageTransport.GinUploadImage(appCtx))
			upload.GET("list", imageTransport.GinListImage(appCtx))
			upload.DELETE("delete/:id", imageTransport.GinDeleteImage(appCtx))
		}

		{
			uploadFood := admin.Group("imagefood")
			uploadFood.POST("create", imageFoodTransport.GinCreateImageFood(appCtx))
			uploadFood.GET("list", imageFoodTransport.GinListImageFood(appCtx))
			uploadFood.DELETE("delete/:id", imageFoodTransport.GinDeleteImageFood(appCtx))
		}

		{
			food := admin.Group("food")
			food.GET("/listfood", foodTransport.GinListFood(appCtx))
			food.POST("/updatefood/:id", foodTransport.GinUpdateFood(appCtx))
			food.POST("/createfood", foodTransport.GinCreateFood(appCtx))
			food.POST("/create-food-with-category/:categoryId", foodTransport.GinCreateFoodAndInfo(appCtx))
			food.DELETE("/deletefood/:id", foodTransport.GinDeleteFood(appCtx))
		}

		{
			user := admin.Group("user")
			user.GET("/list", userTransport.GinListUser(appCtx))
			user.DELETE("/delete/:id", userTransport.GinDeleteUser(appCtx))
			user.PATCH("/update-pass/:id", userTransport.GinUpdateUser(appCtx))
		}

		{
			profile := admin.Group("profile")
			profile.GET("/list", profileTransport.GinListProfile(appCtx))
			profile.PUT("update/:id", profileTransport.GinUpdateProfile(appCtx))
		}

		{
			cart := admin.Group("cart")
			cart.GET("/create", cartTransport.GinCreateCart(appCtx))
			cart.DELETE("/delete", cartTransport.GinDeleteCart(appCtx))
		}

		{
			order := admin.Group("order")
			order.GET("/list", orderTransport.GinListOrder(appCtx))
			order.POST("/create", orderTransport.GinCreateOrder(appCtx))
			order.POST("/update-state/:id", orderTransport.GinUpdateOrderState(appCtx))
		}

		{
			orderDetail := admin.Group("orderdetail")
			orderDetail.GET("/list", orderDetailTransport.GinListOrderDetail(appCtx))
			orderDetail.POST("/create", orderDetailTransport.GinCreateOrderDetail(appCtx))
		}

		{
			category := admin.Group("category")
			category.GET("/list", categoryTransport.GinListCategory(appCtx))
			category.POST("/create", categoryTransport.GinCreateCategory(appCtx))
			category.DELETE("/delete/:id", categoryTransport.GinDeleteCategory(appCtx))
			category.POST("/update/:id", categoryTransport.GinUpdateCategory(appCtx))
		}

		{
			infoFoodCategory := admin.Group("infoFoodCategory")
			infoFoodCategory.GET("/list", infoFoodcategoryTransport.GinListInfoFoodCategory(appCtx))
			infoFoodCategory.POST("/create", infoFoodcategoryTransport.GinCreateInfoFoodCategory(appCtx))
			infoFoodCategory.DELETE("/delete/:id", infoFoodcategoryTransport.GinDeleteInfoFoodCategory(appCtx))
			infoFoodCategory.POST("/update/:id", infoFoodcategoryTransport.GinUpdateInfoFoodCategory(appCtx))
		}

		{
			brand := admin.Group("brand")
			brand.GET("/list", brandTransport.GinListBrand(appCtx))
			brand.POST("/create", brandTransport.GinCreateBrand(appCtx))
			brand.DELETE("/delete/:id", brandTransport.GinDeleteBrand(appCtx))
			brand.POST("/update/:id", brandTransport.GinUpdateBrand(appCtx))
		}

	}
	//user
	{
		user := rt.Group("user")
		user.POST("register", userTransport.GinRegister(appCtx))
		user.POST("authenticate", userTransport.GinLogin(appCtx))
		user.DELETE("delete/:id",
			middleware.RequiredAuth(appCtx),
			userTransport.GinDeleteUser(appCtx),
		)
		user.PATCH("update-pass/:id",
			middleware.RequiredAuth(appCtx),
			userTransport.GinUpdateUser(appCtx),
		)
	}

	{
		profile := rt.Group("profile", middleware.RequiredAuth(appCtx))
		profile.PUT("update/:id", profileTransport.GinUpdateProfile(appCtx))
	}

	{
		rt.POST("payment", middleware.RequiredAuth(appCtx), paymentTransport.GinPayment(appCtx))
	}

	{
		cart := rt.Group("cart", middleware.RequiredAuth(appCtx))
		cart.GET("/list", cartTransport.GinListCart(appCtx))
	}

	{
		category := rt.Group("category")
		category.GET("/list", categoryTransport.GinListCategory(appCtx))
	}

	{
		province := rt.Group("province")
		province.GET("/list", provinceTransport.GinListProvince(appCtx))
	}

	rt.Run()
}
