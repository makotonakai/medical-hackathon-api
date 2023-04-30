package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type Profile struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID string `json:"user_id" param:"user_id"`
	Sex string `json:"sex"`
	IsPregnant bool `json:"is_pregnant"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func (m * Profile) Create() {
	db.Create(&m)
}

func (m *Profile) GetAll() []Profile {
	mh := []Profile{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func (m *Profile) GetByUserId(user_id int) {
	err := db.First(&m,"user_id = ?", user_id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
}
