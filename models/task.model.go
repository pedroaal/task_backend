package models

import (
	"server/config"

	_ "gorm.io/driver/sqlite"
)

func FindTasks(body *[]Task) (err error) {
	err = config.DB.Find(body).Error

	if err != nil {
		return err
	}

	return nil
}

func CreateTask(body *Task) (err error) {
	err = config.DB.Create(body).Error

	if err != nil {
		return err
	}

	return nil
}

func FindTaskById(body *Task, id string) (err error) {
	err = config.DB.Where("id = ?", id).First(body).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateTask(body *Task, id string) (err error) {
	err = config.DB.Save(body).Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteTask(body *Task, id string) (err error) {
	err = config.DB.Where("id = ?", id).Delete(body).Error

	if err != nil {
		return err
	}

	return nil
}
