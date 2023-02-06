package main

import "github.com/labstack/echo/v4"

func main() {
	e := echo.New()
	//router := e.Router()

	e.Logger.Fatal(e.Start(":8080"))
}
