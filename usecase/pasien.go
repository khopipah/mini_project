package usecase

import (
	"errors"
	"fmt"
	"mini_project/model"
	"mini_project/repository/database"
)

func CreatePasien(pasien *model.Pasien) error {

	// check name cannot be empty
	if pasien.Name == "" {
		return errors.New("pasien name cannot be empty")
	}

	// check gender
	if pasien.Gender == "" {
		return errors.New("pasien gender cannot be empty")
	}

	// check Alamat
	if pasien.Alamat == "" {
		return errors.New("pasien Alamat cannot be empty")
	}

	// check telp
	if pasien.Telp == "" {
		return errors.New("pasien telp cannot be empty")
	}

	err := database.CreatePasien(pasien)
	if err != nil {
		return err
	}

	return nil
}

func GetPasien(id uint) (pasien model.Pasien, err error) {
	pasien, err = database.GetPasien(id)
	if err != nil {
		fmt.Println("GetPasien: Error getting docter from database")
		return
	}
	return
}

func GetListPasiens() (pasiens []model.Pasien, err error) {
	pasiens, err = database.GetPasiens()
	if err != nil {
		fmt.Println("GetListPasiens: Error getting pasiens from database")
		return
	}
	return
}

func UpdatePasien(pasien *model.Pasien) (err error) {
	err = database.UpdatePasien(pasien)
	if err != nil {
		fmt.Println("UpdatePasien : Error updating pasien, err: ", err)
		return
	}

	return
}

func DeletePasien(id uint) (err error) {
	pasien := model.Pasien{}
	pasien.ID = id
	err = database.DeletePasien(&pasien)
	if err != nil {
		fmt.Println("DeletePasien : error deleting pasien, err: ", err)
		return
	}

	return
}
