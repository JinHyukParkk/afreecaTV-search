package api

import (
	"net/http"
	"github.com/labstack/echo"
)

func GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
	// return func(c echo.Context) (err error) {
	// 	return '{id="11"}';
	// }
}

func Test(c echo.Context) error {
	u :="{!!!}"
	return c.JSON(http.StatusOK, u)
}

func Test1() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, nil)
	}
}