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

func CreateProfile(c echo.Context) error {
	
	p := model.Profile{}
	
	err := c.Bind(&p)
	if err != nil {
		fmt.Println("Failed to create Profile")
	}

	p.Create()
	return c.JSON(http.StatusCreated, p)
	
}

func GetProfiles(c echo.Context) error {

	p := model.Profile{}
	pList := p.GetAll()
	return c.JSON(http.StatusOK, pList)

}

func GetProfileByUserId(c echo.Context) error {

	p := model.Profile{}
	err := c.Bind(&p)

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}
	p.GetByUserId(user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, p)
}
