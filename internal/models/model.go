package models

import (
	"mime/multipart"
	"smoke-detect/config"
)

type IModel interface {
	UploadFile(file multipart.File, object string) error
	CreateImageData(data map[string]interface{}) error
	GetAllData() []map[string]interface{}
}

type Model struct {
	Config *config.Config
}

func NewModel(cf *config.Config) IModel {
	return &Model{
		Config: cf,
	}
}

func GetHome() string {
	return "Hellooo"
}
