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

func CreateUser(c echo.Context) error {
	
	user := model.User{}
	
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("Failed to create user")
	}

	user.Create()
	return c.JSON(http.StatusCreated, user)
	
}

func GetUsers(c echo.Context) error {

	u := model.User{}
	userList := u.GetAll()
	return c.JSON(http.StatusOK, userList)

}

func GetUserById(c echo.Context) error {

	user := model.User{}
	err := c.Bind(&user)

	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		fmt.Println(err.Error())
	}
	user.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {

	user := model.User{}
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Update()
	return c.JSON(http.StatusOK, user)
}

func DeleteUser(c echo.Context) error {

	user := model.User{}
	
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.Delete()
	return c.JSON(http.StatusNoContent, user)
}
