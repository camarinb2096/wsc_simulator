package players

import (
	"time"

	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Name        string
	Nacionality string
	BirthDate   time.Time
	Age         int
	Position    string
	Number      int
	Photo       string
	FkTeam      int
}
