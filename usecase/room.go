package usecase

import (
	"errors"
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreateRoom(room *model.Room) error {

	// check name cannot be empty
	if room.Name == "" {
		return errors.New("room name cannot be empty")
	}

	// check class
	if room.Class == "" {
		return errors.New("room class cannot be empty")
	}

	// check no
	if room.No == "" {
		return errors.New("room no cannot be empty")
	}

	err := database.CreateRoom(room)
	if err != nil {
		return err
	}

	return nil
}

func GetRoom(id uint) (room model.Room, err error) {
	room, err = database.GetRoom(id)
	if err != nil {
		fmt.Println("GetRoom: Error getting room from database")
		return
	}
	return
}

func GetListRooms() (rooms []model.Room, err error) {
	rooms, err = database.GetRooms()
	if err != nil {
		fmt.Println("GetListRooms: Error getting rooms from database")
		return
	}
	return
}

func UpdateRoom(room *model.Room) (err error) {
	err = database.UpdateRoom(room)
	if err != nil {
		fmt.Println("UpdateRoom : Error updating room, err: ", err)
		return
	}

	return
}

func DeleteRoom(id uint) (err error) {
	room := model.Room{}
	room.ID = id
	err = database.DeleteRoom(&room)
	if err != nil {
		fmt.Println("DeleteRoom : error deleting room, err: ", err)
		return
	}

	return
}
