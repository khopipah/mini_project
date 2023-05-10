package model

import (
	"gorm.io/gorm"
)

type Docter struct {
	gorm.Model
	Name         string `json:"name" form:"name"`
	Gender       string `json:"gender" form:"gender"`
	Telp         string `json:"telp" form:"telp"`
	Alamat       string `json:"alamat" form:"alamat"`
	Id_spesialis string `json:"id_spesialis" form:"id_spesialis"`
}
