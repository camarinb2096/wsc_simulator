package posititons

import "gorm.io/gorm"

type Position struct {
	gorm.Model
	Name      string
	ShortName string
}
