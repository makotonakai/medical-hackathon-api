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

func CreateMedicalHistory(c echo.Context) error {
	
	mh := model.MedicalHistory{}
	
	err := c.Bind(&mh)
	if err != nil {
		fmt.Println("Failed to create MedicalHistory")
	}

	mh.Create()
	return c.JSON(http.StatusCreated, mh)
	
}

func GetMedicalHistories(c echo.Context) error {

	mh := model.MedicalHistory{}
	mhList := mh.GetAll()
	return c.JSON(http.StatusOK, mhList)

}

func GetMedicalHistoriesByUserId(c echo.Context) error {

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}

	mh := model.GetMedicalHistoriesByUserId(user_id)
	return c.JSON(http.StatusOK, mh)
}
