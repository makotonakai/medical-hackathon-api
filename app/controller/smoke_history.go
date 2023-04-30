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

func CreateSmokeHistory(c echo.Context) error {
	
	sh := model.SmokeHistory{}
	
	err := c.Bind(&sh)
	if err != nil {
		fmt.Println("Failed to create SmokeHistory")
	}

	sh.Create()
	return c.JSON(http.StatusCreated, sh)
	
}

func GetSmokeHistories(c echo.Context) error {

	sh := model.SmokeHistory{}
	shList := sh.GetAll()
	return c.JSON(http.StatusOK, shList)

}

func GetSmokeHistoriesByUserId(c echo.Context) error {

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}
	sh := model.GetSmokeHistoriesByUserId(user_id)
	return c.JSON(http.StatusOK, sh)
}
