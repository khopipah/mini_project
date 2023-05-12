package controller

import (
	"mini_project/model"
	"net/http"
	"strconv"

	"mini_project/usecase"

	"github.com/labstack/echo/v4"
)

func GetPasiensController(c echo.Context) error {
	pasiens, e := usecase.GetListPasiens()
	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"pasiens": pasiens,
	})
}

func GetPasienController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pasien, err := usecase.GetPasien(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "data pasien tidak tersedia",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  "success",
		"pasiens": pasien,
	})
}

// create new docter
func CreatePasienController(c echo.Context) error {
	pasien := model.Pasien{}
	c.Bind(&pasien)

	if err := usecase.CreatePasien(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create pasien",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new pasien",
		"pasien":  pasien,
	})
}

// delete docter by id
func DeletePasienController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeletePasien(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete pasien",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pasien tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete pasien",
	})
}

// update docter by id
func UpdatePasienController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pasien := model.Pasien{}
	c.Bind(&pasien)
	pasien.ID = uint(id)
	if err := usecase.UpdatePasien(&pasien); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update pasien",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf pasien tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update pasien",
	})
}
