package database

import (
	"mini_project/config"
	"mini_project/model"
)

func CreateRoom(room *model.Room) error {
	if err := config.DB.Create(room).Error; err != nil {
		return err
	}
	return nil
}

func GetRooms() (rooms []model.Room, err error) {
	if err = config.DB.Find(&rooms).Error; err != nil {
		return
	}
	return
}

func GetRoom(id uint) (room model.Room, err error) {
	room.ID = id
	if err = config.DB.First(&room).Error; err != nil {
		return
	}
	return
}

func UpdateRoom(room *model.Room) error {
	if err := config.DB.Updates(room).Error; err != nil {
		return err
	}
	return nil
}

func DeleteRoom(room *model.Room) error {
	if err := config.DB.Delete(room).Error; err != nil {
		return err
	}
	return nil
}
