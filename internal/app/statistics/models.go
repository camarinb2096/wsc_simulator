package statistics

import "gorm.io/gorm"

type Statistic struct {
	gorm.Model
	FkMatch            uint
	LocalGoals         int
	VisitorGoals       int
	LocalYellowCards   int
	VisitorYellowCards int
	LocalRedCards      int
	VisitorRedCards    int
}
