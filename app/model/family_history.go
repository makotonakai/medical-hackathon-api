package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type FamilyHistory struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID string `json:"user_id" param:"user_id"`
	Description string `json:"description"`
}

func (m * FamilyHistory) Create() {
	db.Create(&m)
}

func (m *FamilyHistory) GetAll() []FamilyHistory {
	mh := []FamilyHistory{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetFamilyHistoriesByUserId(user_id int) []FamilyHistory {
	fh := []FamilyHistory{}
	err := db.Where("user_id = ?", user_id).Find(&fh).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return fh
}
