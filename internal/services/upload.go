package services

import (
	"fmt"
	"mime/multipart"
	"smoke-detect/internal/utils"

	"github.com/gin-gonic/gin"
)

func (service Service) UploadFile(f *multipart.FileHeader) (interface{}, error) {
	blobFile, err := f.Open()
	if err != nil {
		return nil, err
	}
	f.Filename = utils.GenerateFileName()
	imgUrl := fmt.Sprintf("%s%s", service.Config.ImageURL, f.Filename)

	err = service.Model.UploadFile(blobFile, f.Filename)
	if err != nil {
		return nil, err
	}
	data := map[string]interface{}{
		"message": "test_image",
		"imageUrl": imgUrl,
	}
	err = service.Model.CreateImageData(data)
	if err != nil {
		return nil, err
	}
	return gin.H{
		"message":  "Upload Success",
		"imageUrl": imgUrl,
	}, nil
}
