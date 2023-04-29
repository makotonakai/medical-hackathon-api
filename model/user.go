package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
	"github.com/makotonakai/medical-hackathon/database"
)

var db = database.Connect()

type User struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name" param:"name"`
}

func (u *User) Create() {
	db.Create(&u)
}

func (u *User) GetAll() []User {
	users := []User{}
	err := db.Find(&users).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return users
}
 
func (u *User) GetById(id int) {
	err := db.Table("users").Find(&u,"id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
}

func (u *User) Update() {
	db.Save(&u)
}

func (u *User) Delete() {
	u_ := User{}
  db.Delete(&u_, "id = ?", u.ID)
}
