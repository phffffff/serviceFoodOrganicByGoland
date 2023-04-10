package imageTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	hash "go_service_food_organic/component/hasher"
	imageBusiness "go_service_food_organic/module/upload/image/business"
	imageRepo "go_service_food_organic/module/upload/image/repository"
	imageStorage "go_service_food_organic/module/upload/image/storage"
	"net/http"
)

func GinUploadImage(appCtx appContext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		//fileHeader, err := c.FormFile("file")
		//if err != nil {
		//	panic(err)
		//}
		//if err := c.SaveUploadedFile(fileHeader, fmt.Sprintf("static/%s", fileHeader.Filename)); err != nil {
		//	panic(err)
		//}
		//c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(common.Image{
		//	Id:        0,
		//	Url:       fmt.Sprintf("http://localhost:8080/static/%s", fileHeader.Filename),
		//	Width:     400,
		//	Height:    400,
		//	CloudName: "local",
		//	Extension: "png",
		//}))

		//aws 3s
		//
		fileHeader, err := c.FormFile("file")
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")

		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		defer file.Close()

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		db := appCtx.GetMyDBConnection()
		hasher := hash.NewMd5Hash()

		store := imageStorage.NewSqlModel(db)
		repo := imageRepo.NewUploadImageRepo(appCtx.UploadProvider(), store, hasher)
		biz := imageBusiness.NewUploadImageBiz(repo)

		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		img.Mark(false)
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
