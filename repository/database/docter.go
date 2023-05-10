package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateDocter(docter *model.Docter) error {
	if err := config.DB.Create(docter).Error; err != nil {
		return err
	}
	return nil
}

func GetDocters() (docters []model.Docter, err error) {
	if err = config.DB.Find(&docters).Error; err != nil {
		return
	}
	return
}

func GetDocter(id uint) (docter model.Docter, err error) {
	docter.ID = id
	if err = config.DB.First(&docter).Error; err != nil {
		return
	}
	return
}

func UpdateDocter(docter *model.Docter) error {
	if err := config.DB.Updates(docter).Error; err != nil {
		return err
	}
	return nil
}

func DeleteDocter(docter *model.Docter) error {
	if err := config.DB.Delete(docter).Error; err != nil {
		return err
	}
	return nil
}
