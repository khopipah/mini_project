package model

import (
	"gorm.io/gorm"
)

type Room struct {
	gorm.Model
	Name  string `json:"name" form:"name"`
	Class string `json:"class" form:"class"`
	No    string `json:"no" form:"no"`
}
