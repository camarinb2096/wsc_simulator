package phases

import "gorm.io/gorm"

type Phase struct {
	gorm.Model
	Name string
}
