package controllers

import (
	"smoke-detect/config"
	"smoke-detect/internal/services"

	"github.com/gin-gonic/gin"
)

type IController interface {
	GetHome(c *gin.Context) 
	UploadFile(c *gin.Context)
	GetAllDetectedSmoker(c *gin.Context)
}

type Controller struct {
	Service services.IService
}

func NewController(cf *config.Config) IController {
	return &Controller{
		Service: services.NewService(cf),
	}
}

func (controller Controller) GetHome(c *gin.Context) {
	home := controller.Service.Home()
	c.JSON(200, gin.H{
		"message": "Home Page",
		"data": home,
	})
}
