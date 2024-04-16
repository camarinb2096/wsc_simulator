package teams

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	Name      string
	FlagPhoto string
}
