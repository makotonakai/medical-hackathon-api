package router

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/makotonakai/medical-hackathon/app/controller"
)

func Initialize() *echo.Echo {

	e := echo.New()
	e.Use(middleware.Logger())
  e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	api := e.Group("/api")

	api.POST("/users/new", controller.CreateUser)
	api.GET("/users", controller.GetUsers)
	api.GET("/users/:id", controller.GetUserById)
	api.PUT("/users/:id", controller.UpdateUser)
	api.DELETE("/users/:id", controller.DeleteUser)

	api.POST("/clinics/new", controller.CreateClinic)
	api.GET("/clinics", controller.GetClinics)
	api.GET("/clinics/:id", controller.GetClinicById)
	api.PUT("/clinics/:id", controller.UpdateClinic)
	api.DELETE("/clinics/:id", controller.DeleteClinic)

	api.POST("/mypage/new", controller.CreateMyPage)
	api.GET("/mypage/:user_id", controller.GetMyPageByUserId)

	return e

}
