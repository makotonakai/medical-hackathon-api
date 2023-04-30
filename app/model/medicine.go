package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type Medicine struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID string `json:"user_id" param:"user_id"`
	Description string `json:"description"`
}

func (m * Medicine) Create() {
	db.Create(&m)
}

func (m *Medicine) GetAll() []Medicine {
	mh := []Medicine{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetMedicinesByUserId(user_id int) []Medicine {
	m := []Medicine{}
	err := db.Where("user_id = ?", user_id).Find(&m).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return m
}
