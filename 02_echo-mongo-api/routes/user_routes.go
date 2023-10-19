package routes

import (
	"echo-mongo-api/controllers"

	"github.com/labstack/echo"
)

func UserRoute(e *echo.Echo) {
	// All routes related to users goes here
	e.POST("/user", controllers.CreateUser)
	e.GET("/user/:userId", controllers.GetAUser)
}
