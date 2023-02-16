package app

import (
	"awesomeProject2/utils"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Hello(c echo.Context) (err error) {
	return utils.SendResponse(c, utils.ResponseMessage{
		Code:    http.StatusOK,
		Message: "data response",
		Data:    "“Hello Go developers",
	})
}

func CheckPalindrom(c echo.Context) (err error) {
	str := c.FormValue("str")

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			print, _ := fmt.Println("bukan palindrome")
			fmt.Println(print)
			return
		}
	}
	print, _ := fmt.Println("ini palindrome")
	fmt.Println(print)
	return
}
