package routers

import (
	"github.com/gin-gonic/gin"
	"smoke-detect/config"

)

func InitRouter(cf *config.Config) *gin.Engine {
	r := gin.Default()
	handlers := GetHandlers(cf)
	for _, handler := range handlers {
		r.Handle(handler.Method, handler.Path, handler.Function)
	}
	return r
}
