package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type DrinkHistory struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID int `json:"user_id"  param:"user_id"`
	Description string `json:"description"`
}

func (m * DrinkHistory) Create() {
	db.Create(&m)
}

func (m *DrinkHistory) GetAll() []DrinkHistory {
	mh := []DrinkHistory{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetDrinkHistoriesByUserUserId(user_id int) []DrinkHistory {
	dh := []DrinkHistory{}
	err := db.Table("drink_histories").Where("user_id = ?", user_id).Find(&dh).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return dh
}
