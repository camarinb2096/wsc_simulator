package players

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Upload Controller
		Get    Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Upload: func(c *gin.Context, s Services) {
			UploadData(c, s)
		},
		Get: func(c *gin.Context, s Services) {
			GetPlayers(c, s)
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

	err = s.Create(file)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "players created successfully",
	})

}

func GetPlayers(c *gin.Context, s Services) {
	fkTeam := c.DefaultQuery("team", "0")

	fkTeamInt, err := strconv.Atoi(fkTeam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid fk_team"})
		return
	}
	players, err := s.Get(fkTeamInt)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(players) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "players not found"})
		return
	}
	c.JSON(200, gin.H{
		"message": "players retrieved successfully",
		"players": players,
	})
}
