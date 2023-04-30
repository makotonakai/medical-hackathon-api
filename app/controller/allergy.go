package controller

import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/makotonakai/medical-hackathon/app/model"
)

//----------
// Handlers
//----------

func CreateAllergy(c echo.Context) error {
	
	Allergy := model.Allergy{}
	
	err := c.Bind(&Allergy)
	if err != nil {
		fmt.Println("Failed to create Allergy")
	}

	Allergy.Create()
	return c.JSON(http.StatusCreated, Allergy)
	
}

func GetAllergies(c echo.Context) error {

	Allergy := model.Allergy{}
	AllergyList := Allergy.GetAll()
	return c.JSON(http.StatusOK, AllergyList)

}

func GetAllergiesByUserId(c echo.Context) error {

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}

	Allergies := model.GetAllergiesByUserId(user_id)
	return c.JSON(http.StatusOK, Allergies)
}
