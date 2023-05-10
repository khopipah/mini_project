package model

import (
	"gorm.io/gorm"
)

type Pasien struct {
	gorm.Model
	Name     string `json:"name" form:"name"`
	Gender   string `json:"gender" form:"gender"`
	Telp     string `json:"telp" form:"telp"`
	Alamat   string `json:"alamat" form:"alamat"`
	Penyakit string `json:"penyakit" form:"penyakit"`
}
