package main

import (
	"log"
	"net/http"
	"nomni/utils/validator"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

func main() {

	e := echo.New()

	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})

	e.POST("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	r := echoswagger.New(e, "docs", &echoswagger.Info{
		Title:       "Sample Fruit API",
		Description: "This is docs for fruit service",
		Version:     "1.0.0",
	})
	r.AddSecurityAPIKey("Authorization", "JWT token", echoswagger.SecurityInHeader)
	r.SetUI(echoswagger.UISetting{
		HideTop: true,
	})

	FruitApiController{}.Init(r.Group("fruits", "/fruits"))
	e.Validator = validator.New()
	if err := e.Start(":8080"); err != nil {
		log.Println(err)
	}
}
