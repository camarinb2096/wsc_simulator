package championship

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Start   Controller
		Restart Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Start: func(c *gin.Context, s Services) {
			StartChampionship(c, s)

		},
		Restart: func(c *gin.Context, s Services) {
			RestartChampionship(c, s)
		},
	}
}

func StartChampionship(c *gin.Context, s Services) {
	s.PlayChampionship()

	c.JSON(200, gin.H{
		"message": "Playing Championship",
	})

}

func RestartChampionship(c *gin.Context, s Services) {
	err := s.RestartChampionship()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error Restarting Championship",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Championship Restarted",
	})

}
