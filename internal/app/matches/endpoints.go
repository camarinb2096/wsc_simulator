package matches

import (
	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Get           Controller
		GetStatistics Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Get: func(c *gin.Context, s Services) {
			GetMatches(c, s)
		},
		GetStatistics: func(c *gin.Context, s Services) {
			GetStatistics(c, s)
		},
	}
}

func GetMatches(c *gin.Context, s Services) {
	response, err := s.GetMatches()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting matches",
		})
		return
	}
	if response.Total == 0 {
		c.JSON(404, gin.H{
			"message": "no matches found",
		})
		return
	}
	c.JSON(200, response)
}

func GetStatistics(c *gin.Context, s Services) {
	response, err := s.GetStatistics()

	if err != nil {
		c.JSON(500, gin.H{
			"message": "error getting statistics",
		})
		return
	}

	c.JSON(200, response)
}
