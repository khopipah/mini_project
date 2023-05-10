package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreatePasien(pasien *model.Pasien) error {
	if err := config.DB.Create(pasien).Error; err != nil {
		return err
	}
	return nil
}

func GetPasiens() (pasiens []model.Pasien, err error) {
	if err = config.DB.Find(&pasiens).Error; err != nil {
		return
	}
	return
}

func GetPasien(id uint) (pasien model.Pasien, err error) {
	pasien.ID = id
	if err = config.DB.First(&pasien).Error; err != nil {
		return
	}
	return
}

func UpdatePasien(pasien *model.Pasien) error {
	if err := config.DB.Updates(pasien).Error; err != nil {
		return err
	}
	return nil
}

func DeletePasien(pasien *model.Pasien) error {
	if err := config.DB.Delete(pasien).Error; err != nil {
		return err
	}
	return nil
}
