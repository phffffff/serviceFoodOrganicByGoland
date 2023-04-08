package uploadTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	uploadBusiness "go_service_food_organic/module/upload/business"
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

		biz := uploadBusiness.NewUploadImageBiz(appCtx.UploadProvider(), nil)
		img, err := biz.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
