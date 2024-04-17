package teams

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	Controller func(c *gin.Context, s Services)

	Endpoints struct {
		Upload      Controller
		Get         Controller
		GetChampion Controller
	}
)

func NewEndpoints(s Services) *Endpoints {
	return &Endpoints{
		Upload: func(c *gin.Context, s Services) {
			UploadData(c, s)
		},
		Get: func(c *gin.Context, s Services) {
			GetTeams(c, s)
		},
		GetChampion: func(c *gin.Context, s Services) {
			GetChampionTeam(c, s)
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

func GetTeams(c *gin.Context, s Services) {
	teams, err := s.Get()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting teams",
		})
		return
	}
	if len(teams) == 0 {
		c.JSON(404, gin.H{
			"message": "No teams found",
		})
		return
	}
	c.JSON(200, teams)
}

func GetChampionTeam(c *gin.Context, s Services) Team {
	team, err := s.GetChampionTeam()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting champion team",
		})
		return Team{}
	}
	c.JSON(200, team)
	return team
}
