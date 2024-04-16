package championship

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Start Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Start: StartChampionship,
	}
}

func StartChampionship(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "StartChampionship",
	})

}
