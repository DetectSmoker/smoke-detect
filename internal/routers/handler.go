package routers

import (
	"smoke-detect/config"
	"smoke-detect/internal/controllers"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Path string
	Method   string
	Function gin.HandlerFunc
}

type Handlers []Handler 

func GetHandlers(cf *config.Config) Handlers {
	controllers := controllers.NewController(cf)
	handlers := Handlers{
		{
			Path: "/",
			Method: "GET",
			Function: controllers.GetHome,
		},
		{
			Path: "/upload/file",
			Method: "POST",
			Function: controllers.UploadFile,
		},
		{
			Path: "/detected-smoker",
			Method: "GET",
			Function: controllers.GetAllDetectedSmoker,
		},
	}
	return handlers
}
