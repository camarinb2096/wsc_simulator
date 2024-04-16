package players

import (
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name        string
	Nacionality string
	BirthDate   string
	Position    string
	Number      int
	Photo       string
	FkTeam      int
}
