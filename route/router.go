package route

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo"
	echoMw "github.com/labstack/echo/middleware"
	"github.com/JinHyukParkk/randomProject/api"
)

type (
	Host struct {
		Echo *echo.Echo
	}
)

func Init() *echo.Echo {
	// Hosts
	hosts := map[string]*Host{}

	//-----
	// API
	//-----

	searchapi := echo.New()
	searchapi.Use(echoMw.Logger())
	searchapi.Use(echoMw.Recover())

	//SubDomain
	hosts["searchapi.afreecatv.com:8080"] = &Host{searchapi}

	searchapi.GET("/test", api.Test)
	searchapi.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "API")
	})

	// e.Debug()

	e := echo.New()
	// Set Bundle MiddleWare
	e.Use(echoMw.Logger())
	e.Use(echoMw.Recover())
	e.Use(echoMw.Gzip())
	e.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	//HTTP Error Handler
	e.HTTPErrorHandler = customHTTPErrorHandler


	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	// e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := e.Group("/api/v1")
	{
		// v1.POST("/search", api.PostMember())
		v1.GET("/search", api.GetList)
		// v1.GET("/members/:id", api.GetMember())
	}

	e.GET("/test", api.Test)
	e.GET("/test1", api.Test1())
	e.GET("/ping", func(c echo.Context) error {
        return c.String(200, "test")
	})
	
	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	errorPage := fmt.Sprintf("errorHtml/%d.html", code)
	if err := c.File(errorPage); err != nil {
		c.Logger().Error(err)
	}
	// c.Logger().Error(err)
}
