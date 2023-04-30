package model

import (
	"fmt"
	"time"
	"errors"
	"gorm.io/gorm"
)

type Reservation struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	UserID int`json:"user_id" param:"user_id"`
	ClinicID int `json:"clinic_id" param:"clinic_id"`
	StartTime time.Time `json:"start_time"`
	Symptoms []string
	Description string `json:"description"`
} 

func (m * Reservation) Create() {
	db.Create(&m)
}

func (m *Reservation) GetAll() []Reservation {
	mh := []Reservation{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetReservationByUserId(user_id int) []Reservation {
	sh := []Reservation{}
	err := db.Table("reservations").Where("user_id = ?", user_id).Find(&sh).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
	return sh
}
