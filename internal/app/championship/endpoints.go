package championship

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Start Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Start: func(c *gin.Context, s Services) {
			StartChampionship(c, s)
		},
	}
}

func StartChampionship(c *gin.Context, s Services) {
	s.PlayChampionship()

	c.JSON(200, gin.H{
		"message": "Playing Championship",
	})

}
