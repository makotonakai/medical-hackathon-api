package model

import (
	"fmt"
	"errors"
	"gorm.io/gorm"
)

type Clinic struct {
	ID int `json:"id" param:"id" gorm:"primaryKey;autoIncrement"`
	Name string `json:"name"`
	PostalCode string `json:"postal_code"`
	Address string `json:"address"`
}

func (c *Clinic) Create() {
	db.Create(&c)
}

func (c *Clinic) GetAll() []Clinic {
	clinics := []Clinic{}
	err := db.Find(&clinics).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return clinics
}
 
func (c *Clinic) GetById(id int) {
	err := db.Table("clinics").Find(&c,"id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("No user found")
	}
}

func (c *Clinic) Update() {
	db.Save(&c)
}

func (c *Clinic) Delete() {
	c_ := Clinic{}
  db.Delete(&c_, "id = ?", c.ID)
}
