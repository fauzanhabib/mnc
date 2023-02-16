package app

import "github.com/labstack/echo/v4"

func AssignRouting(e *echo.Echo) {
	e.GET("/", Hello)
	e.GET("/polindrom", CheckPalindrom)

}
