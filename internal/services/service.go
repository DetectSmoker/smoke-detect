package services

import (
	"mime/multipart"
	"smoke-detect/config"
	"smoke-detect/internal/models"
)

type IService interface {
	Home() interface{}
	UploadFile(f *multipart.FileHeader) (interface{}, error)
	GetAllDetectedSmoker() (interface{}, error)
}

type Service struct {
	Model  models.IModel
	Config *config.Config
}

func NewService(cf *config.Config) IService {
	return &Service{
		Model: models.NewModel(cf),
		Config: cf,
	}
}
