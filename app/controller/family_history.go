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

func CreateFamilyHistory(c echo.Context) error {
	
	fh := model.FamilyHistory{}
	
	err := c.Bind(&fh)
	if err != nil {
		fmt.Println("Failed to create FamilyHistory")
	}

	fh.Create()
	return c.JSON(http.StatusCreated, fh)
	
}

func GetFamilyHistories(c echo.Context) error {

	fh := model.FamilyHistory{}
	fhList := fh.GetAll()
	return c.JSON(http.StatusOK, fhList)

}

func GetFamilyHistoriesByUserId(c echo.Context) error {

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}
	

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	fh := model.GetFamilyHistoriesByUserId(user_id)
	return c.JSON(http.StatusOK, fh)
}
