package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type Allergy struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID int `json:"user_id" param:"user_id"`
	Description string `json:"description"`
}

func (m * Allergy) Create() {
	db.Create(&m)
}

func (m *Allergy) GetAll() []Allergy {
	mh := []Allergy{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetAllergiesByUserId(user_id int) []Allergy {
	a := []Allergy{}
	err := db.Where("user_id = ?", user_id).Find(&a).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return a
}
