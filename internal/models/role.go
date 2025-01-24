package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name			string `gorm:"size:50;not null;unique" json:"name"`
	Description		string 	`gorm:"size:255;not null" json:"description"`
}

func CreateRole(Role *Role) (err error) {
	err = DB.Create(Role).Error

	if err != nil {
		return err
	}
	return nil
}

func GetRoles(Role *[]Role)(err error) {
	err = DB.Find(Role).Error

	if err != nil {
		return err
	}
	return nil
}

func GetRole(Role *Role, id int) (err error) {
	err = DB.Where("id = ?", id).First(Role).Error

	if err != nil {
		return err
	}
	return nil
}

func UpdateRole(Role *Role) (err error) {
	DB.Save(Role)
	return nil
}