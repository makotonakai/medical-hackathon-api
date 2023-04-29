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

func CreateClinic(c echo.Context) error {
	
	clinic := model.Clinic{}
	
	err := c.Bind(&clinic)
	if err != nil {
		fmt.Println("Failed to create Clinic")
	}

	clinic.Create()
	return c.JSON(http.StatusCreated, clinic)
	
}

func GetClinics(c echo.Context) error {

	clinic := model.Clinic{}
	clinicList := clinic.GetAll()
	return c.JSON(http.StatusOK, clinicList)

}

func GetClinicById(c echo.Context) error {

	clinic := model.Clinic{}
	err := c.Bind(&clinic)

	id_ := c.Param("id")
	id, err := strconv.Atoi(id_)
	if err != nil {
		fmt.Println(err.Error())
	}
	clinic.GetById(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, clinic)
}

func UpdateClinic(c echo.Context) error {

	clinic := model.Clinic{}
	err := c.Bind(&clinic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	clinic.Update()
	return c.JSON(http.StatusOK, clinic)
}

func DeleteClinic(c echo.Context) error {

	clinic := model.Clinic{}
	
	err := c.Bind(&clinic)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	clinic.Delete()
	return c.JSON(http.StatusNoContent, clinic)
}
