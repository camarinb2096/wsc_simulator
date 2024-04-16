package players

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Upload Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Upload: UploadData,
	}
}

func UploadData(c *gin.Context) {

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filePath := "internal/app/data/" + file.Filename
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "GetChampionship",
	})

}
