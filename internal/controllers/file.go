package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (controller Controller) UploadFile(c *gin.Context) {
	f, err := c.FormFile("file_input")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := controller.Service.UploadFile(f)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.JSON(200, res)
}

func (controller Controller) GetAllDetectedSmoker(c *gin.Context) {
	res, err := controller.Service.GetAllDetectedSmoker()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, res)
}
