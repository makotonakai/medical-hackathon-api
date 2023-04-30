package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type SmokeHistory struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID string `json:"user_id" param:"user_id"`
	Description string `json:"description"`
}

func (m * SmokeHistory) Create() {
	db.Create(&m)
}

func (m *SmokeHistory) GetAll() []SmokeHistory {
	mh := []SmokeHistory{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetSmokeHistoriesByUserId(user_id int) []SmokeHistory {
	sh := []SmokeHistory{}
	err := db.Where("user_id = ?", user_id).Find(&sh).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return sh
}
