package usecase

import (
	"errors"
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

type DocterUsecase interface {
	CreateDocter(docter *model.Docter) error
	GetDocter(id uint) (docter model.Docter, err error)
	GetListDocter() (docters []model.Docter, err error)
	UpdateDocter(docter *model.Docter) (err error)
	DeleteDocter(id uint) (err error)
}

func CreateDocter(docter *model.Docter) error {

	// check name cannot be empty
	if docter.Name == "" {
		return errors.New("docter name cannot be empty")
	}

	// check gender
	if docter.Gender == "" {
		return errors.New("docter gender cannot be empty")
	}

	// check Alamat
	if docter.Alamat == "" {
		return errors.New("docter Alamat cannot be empty")
	}

	// check id_spesialis
	if docter.Id_spesialis == "" {
		return errors.New("docter id_spesialis cannot be empty")
	}

	// check telp
	if docter.Telp == "" {
		return errors.New("docter telp cannot be empty")
	}

	err := database.CreateDocter(docter)
	if err != nil {
		return err
	}

	return nil
}

func GetDocter(id uint) (docter model.Docter, err error) {
	docter, err = database.GetDocter(id)
	if err != nil {
		fmt.Println("GetDocter: Error getting docter from database")
		return
	}
	return
}

func GetListDocters() (docters []model.Docter, err error) {
	docters, err = database.GetDocters()
	if err != nil {
		fmt.Println("GetListDocters: Error getting docters from database")
		return
	}
	return
}

func UpdateDocter(docter *model.Docter) (err error) {
	err = database.UpdateDocter(docter)
	if err != nil {
		fmt.Println("UpdateDocter : Error updating docter, err: ", err)
		return
	}

	return
}

func DeleteDocter(id uint) (err error) {
	docter := model.Docter{}
	docter.ID = id
	err = database.DeleteDocter(&docter)
	if err != nil {
		fmt.Println("DeleteDocter : error deleting docter, err: ", err)
		return
	}

	return
}
