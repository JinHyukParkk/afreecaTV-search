package api

import (
	"github.com/labstack/echo"
)

func GetList() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return '{id="11"}';
	}
}