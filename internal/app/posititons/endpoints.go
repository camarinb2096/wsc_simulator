package positions

import "github.com/gin-gonic/gin"

type (
	Controller func(c *gin.Context)

	Endpoints struct {
		Get Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Get: GetPositions,
	}
}

func GetPositions(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetPositions",
	})
}
