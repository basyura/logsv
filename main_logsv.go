package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", log)
	e.POST("/", log)

	e.Logger.Fatal(e.Start(":5678"))
}

func log(c echo.Context) error {

	fmt.Println("\n########################################")

	fmt.Println("URLParam ------------")
	for k, v := range c.QueryParams() {
		fmt.Println(" ", k, "=", v)
	}

	for _, name := range c.ParamNames() {
		fmt.Println(name)
	}

	fmt.Println("FORMParam -----------")
	params, err := c.FormParams()
	if err != nil {
		fmt.Println(err)
	} else {
		for k, v := range params {
			fmt.Println(" ", k, "=", v)
		}
	}

	return nil
}
