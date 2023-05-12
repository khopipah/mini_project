package controller

import (
	"mini_project/model"
	"net/http"
	"strconv"

	"mini_project/usecase"

	"github.com/labstack/echo/v4"
)

func GetDoctersController(c echo.Context) error {
	docters, e := usecase.GetListDocters()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"docters": docters,
	})
}

func GetDocterController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	docter, err := usecase.GetDocter(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "data docter tidak tersedia",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"docters": docter,
	})
}

// create new docter
func CreateDocterController(c echo.Context) error {
	docter := model.Docter{}
	c.Bind(&docter)

	if err := usecase.CreateDocter(&docter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create docter",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new docter",
		"docter":  docter,
	})
}

// delete docter by id
func DeleteDocterController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteDocter(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete docter",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf docter tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete docter",
	})
}

// update docter by id
func UpdateDocterController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	docter := model.Docter{}
	c.Bind(&docter)
	docter.ID = uint(id)
	if err := usecase.UpdateDocter(&docter); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update docter",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf docter tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update docter",
	})
}
