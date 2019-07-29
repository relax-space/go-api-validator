package main

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/pangpanglabs/echoswagger"
)

type FruitApiController struct {
}

type Fruit struct {
	Name string `json:"name" valid:"length(0|5)"`
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
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	for _, v1 := range v {
		_, err := govalidator.ValidateStruct(v1)
		if err != nil {
			return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
		}
	}

	return ReturnApiSucc(c, http.StatusOK, v)
}

func (FruitApiController) Struct(c echo.Context) error {
	var v Fruit
	if err := c.Bind(&v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}
	if err := c.Validate(v); err != nil {
		return ReturnApiFail(c, http.StatusBadRequest, ApiErrorParameter, err)
	}

	return ReturnApiSucc(c, http.StatusOK, v)
}
