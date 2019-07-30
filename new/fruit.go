package main

import (
	"net/http"
	"nomni/utils/api"

	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type FruitApiController struct {
}

type Fruit struct {
	Name  string `json:"name" validate:"lte=5"`
	Color string `json:"color" validate:"gte=2"`
}

// localhost:8080/docs
func (d FruitApiController) Init(g echoswagger.ApiGroup) {
	g.SetSecurity("Authorization")

	g.POST("/struct", d.Struct).
		AddParamBody(Fruit{}, "fruit", "new fruit", true)
	g.POST("/slice", d.Slice).AddParamBody([]Fruit{}, "fruit", "new fruit list", true)
}

func (FruitApiController) Slice(c echo.Context) error {
	var v []Fruit
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
	}
	for _, v1 := range v {
		err := c.Validate(v1)
		if err != nil {
			return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
		}
	}

	return ReturnApiSucc(c, http.StatusOK, v)
}

func (FruitApiController) Struct(c echo.Context) error {
	var v Fruit
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
	}
	if err := c.Validate(v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, api.ParameterParsingError(err))
	}

	return ReturnApiSucc(c, http.StatusOK, v)
}
