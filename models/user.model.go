package models

import (
	"server/config"

	_ "gorm.io/driver/sqlite"
)

func FindUsers(body *[]User) (err error) {
	err = config.DB.Find(body).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateUser(body *User) (err error) {
	err = config.DB.Create(body).Error

	if err != nil {
		return err
	}

	return nil
}

func FindUserById(body *User, id string) (err error) {
	err = config.DB.Where("id = ?", id).First(body).Error

	if err != nil {
		return err
	}

	return nil
}

func FindUserByMail(body *User) (err error) {
	err = config.DB.Where("mail = ?", body.Mail).First(body).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(body *User, id string) (err error) {
	err = config.DB.Save(body).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(body *User, id string) (err error) {
	err = config.DB.Where("id = ?", id).Delete(body).Error

	if err != nil {
		return err
	}

	return nil
}
