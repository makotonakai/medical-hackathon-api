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

func CreateMyPage(c echo.Context) error {
	
	mp := model.MyPage{}
	
	err := c.Bind(&mp)
	if err != nil {
		fmt.Println("Failed to create MyPage")
	}

	mp.Create()
	return c.JSON(http.StatusCreated, mp)
	
}

func GetMyPages(c echo.Context) error {

	mp := model.MyPage{}
	mpList := mp.GetAll()
	return c.JSON(http.StatusOK, mpList)

}

func GetMyPageByUserId(c echo.Context) error {

	mp := model.MyPage{}
	err := c.Bind(&mp)

	ui := c.Param("user_id")
	user_id, err := strconv.Atoi(ui)
	if err != nil {
		fmt.Println(err.Error())
	}
	mp = model.GetMyPageByUserId(user_id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, mp)
}
