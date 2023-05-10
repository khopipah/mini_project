package controller

import (
	"mini_project/model"
	"net/http"
	"strconv"

	"mini_project/usecase"

	"github.com/labstack/echo/v4"
)

func GetRoomcontroller(c echo.Context) error {
	rooms, e := usecase.GetListRooms()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"rooms":  rooms,
	})
}

func GetRoomController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	room, err := usecase.GetRoom(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get room",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"rooms":  room,
	})
}

// create new room
func CreateRoomController(c echo.Context) error {
	room := model.Room{}
	c.Bind(&room)

	if err := usecase.CreateRoom(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create room",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new room",
		"room":    room,
	})
}

// delete room by id
func DeleteRoomController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteRoom(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete room",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf room tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete room",
	})
}

// update room by id
func UpdateRoomController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	room := model.Room{}
	c.Bind(&room)
	room.ID = uint(id)
	if err := usecase.UpdateRoom(&room); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update room",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf room tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update room",
	})
}
