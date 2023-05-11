package imageTransport

import (
	"github.com/gin-gonic/gin"
	"go_service_food_organic/common"
	appContext "go_service_food_organic/component/app_context"
	hash "go_service_food_organic/component/hasher"
	"go_service_food_organic/module/image/business"
	"go_service_food_organic/module/image/model"
	"go_service_food_organic/module/image/repository"
	"go_service_food_organic/module/image/storage"
	"mime/multipart"
	"net/http"
	"sync"
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

		form, err := c.MultipartForm()
		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		fileHeaders := form.File["file"]
		if len(fileHeaders) == 0 {
			panic(common.ErrInvalidRequest(nil))
		}

		typeImage := c.DefaultPostForm("type", "img")

		var wg sync.WaitGroup

		db := appCtx.GetMyDBConnection()
		secretSalt := appCtx.GetSecretSaltHashImage()
		hasher := hash.NewMd5Hash(secretSalt)

		req := c.MustGet(common.CurrentUser).(common.Requester)

		store := imageStorage.NewSqlModel(db)
		repo := imageRepo.NewUploadImageRepo(appCtx.UploadProvider(), store, hasher, req)
		biz := imageBusiness.NewUploadImageBiz(repo)

		errsInfo := make(chan *imageModel.ErrorInfo, len(fileHeaders))

		for _, fileHeader := range fileHeaders {
			wg.Add(1)
			go func(fileHeader *multipart.FileHeader) {
				defer wg.Done()

				file, err := fileHeader.Open()
				if err != nil {
					newErrInfo := imageModel.ErrorInfo{
						FileName: fileHeader.Filename,
						ImgInfo:  nil,
						ErrInfo:  err,
					}
					errsInfo <- &newErrInfo
					return
					//panic(common.ErrInvalidRequest(err))
				}
				defer file.Close()
				dataBytes := make([]byte, fileHeader.Size)
				if _, err := file.Read(dataBytes); err != nil {
					newErrInfo := imageModel.ErrorInfo{
						FileName: fileHeader.Filename,
						ImgInfo:  nil,
						ErrInfo:  err,
					}
					errsInfo <- &newErrInfo
					return
					//panic(common.ErrInvalidRequest(err))
				}
				img, err := biz.Upload(c.Request.Context(), dataBytes, typeImage, fileHeader.Filename)
				if err != nil {
					newErrInfo := imageModel.ErrorInfo{
						FileName: fileHeader.Filename,
						ImgInfo:  nil,
						ErrInfo:  err,
					}
					errsInfo <- &newErrInfo
					return
					//panic(err)
				}
				newErrInfo := imageModel.ErrorInfo{
					FileName: fileHeader.Filename,
					ImgInfo:  img,
					ErrInfo:  nil,
				}
				errsInfo <- &newErrInfo
				return
			}(fileHeader)
		}
		go func() {
			wg.Wait()
			close(errsInfo)
		}()

		var errInfos []*imageModel.ErrorInfo
		for errInfo := range errsInfo {
			if errInfo.ImgInfo != nil {
				errInfo.ImgInfo.Mark(false)
			}
			errInfos = append(errInfos, errInfo)
		}

		c.IndentedJSON(http.StatusOK, common.SimpleSuccessResponse(errInfos))
	}
}
