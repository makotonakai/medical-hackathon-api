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

func CreateMedicine(c echo.Context) error {
	
	m := model.Medicine{}
	
	err := c.Bind(&m)
	if err != nil {
		fmt.Println("Failed to create Medicine")
	}

	m.Create()
	return c.JSON(http.StatusCreated, m)
	
}

func GetMedicines(c echo.Context) error {

	m := model.Medicine{}
	mList := m.GetAll()
	return c.JSON(http.StatusOK, mList)

}

func GetMedicineByUserId(c echo.Context) error {

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}

	m := model.GetMedicinesByUserId(user_id)

	return c.JSON(http.StatusOK, m)
}
