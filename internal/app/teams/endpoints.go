package teams

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Upload Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Upload: func(c *gin.Context, s Services) {
			UploadData(c, s)
		},
	}
}

func UploadData(c *gin.Context, s Services) {

	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	file, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	teams, err := s.Create(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "teams created successfully",
		"teams":   teams,
	})
}
