package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type MedicalHistory struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID int `json:"user_id" param:"user_id"`
	Disease string `json:"disease"`
	Description string `json:"description"`
}

func (m * MedicalHistory) Create() {
	db.Create(&m)
}

func (m *MedicalHistory) GetAll() []MedicalHistory {
	mh := []MedicalHistory{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetMedicalHistoriesByUserId(user_id int) []MedicalHistory {
	mh := []MedicalHistory{}
	err := db.Table("medical_histories").Where("user_id = ?", user_id).Find(&mh).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return mh
}
