package model

import (
	"fmt"
)

type MyPage struct {
	ID int `json:"id"`
	UserID int `json:"user_id" param:"user_id"`
	Profile Profile `json:"profile"`
	MedicalHistories []MedicalHistory `json:"medical_histories"`
	Medicines []Medicine `json:"medicines"`
	Allergies []Allergy `json:"allergies"`
	FamilyHistories []FamilyHistory `json:"family_histories"`
	DrinkHistories []DrinkHistory `json:"drink_histories"`
	SmokeHistories []SmokeHistory `json:"smoke_histories"`
}

func (mp * MyPage) Create() {
	db.Create(&mp.Profile)
	db.Create(&mp.MedicalHistories)
	db.Create(&mp.Medicines)
	db.Create(&mp.Allergies)
	db.Create(&mp.FamilyHistories)
	db.Create(&mp.DrinkHistories)
	db.Create(&mp.SmokeHistories)

}

func (mp *MyPage) GetAll() []MyPage {
	mh := []MyPage{}
	err := db.Find(&mh).Error
	if err != nil {
		fmt.Println("Error found")
	}
	return mh
}
 
func GetMyPageByUserId(user_id int) MyPage {

	mp := MyPage{}
	mp.ID = user_id
	mp.UserID = user_id

	p := Profile{}
	p.GetByUserId(user_id)
	mp.Profile = p

	mh := GetMedicalHistoriesByUserId(user_id)
	mp.MedicalHistories = mh

	m := GetMedicinesByUserId(user_id)
	mp.Medicines = m

	a := GetAllergiesByUserId(user_id)
	mp.Allergies = a

	fh := GetFamilyHistoriesByUserId(user_id)
	mp.FamilyHistories = fh 

	dh := GetDrinkHistoriesByUserUserId(user_id)
	mp.DrinkHistories = dh

	sm := GetSmokeHistoriesByUserId(user_id)
	mp.SmokeHistories = sm

	return mp
}


