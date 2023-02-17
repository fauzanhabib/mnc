package main

import (
	"awesomeProject2/app"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/tylerb/graceful"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	srv := &graceful.Server{
		Timeout: 10 * time.Second,
		Server: &http.Server{
			Addr:    ":8888",
			Handler: e,
		},
	}
	app.AssignRouting(e)
	srv.ListenAndServe()
	fmt.Println("servcer Run")
}
