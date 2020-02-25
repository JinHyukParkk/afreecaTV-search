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
		return c.String(http.StatusOK, "searchAPI")
	})

	// e.Debug()

	test := echo.New()
	// Set Bundle MiddleWare
	test.Use(echoMw.Logger())
	test.Use(echoMw.Recover())
	test.Use(echoMw.Gzip())
	test.Use(echoMw.CORSWithConfig(echoMw.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
	}))

	//SubDomain
	hosts["test.afreecatv.com:8080"] = &Host{test}

	//HTTP Error Handler
	test.HTTPErrorHandler = customHTTPErrorHandler


	// e.SetHTTPErrorHandler(handler.JSONHTTPErrorHandler)

	// Set Custom MiddleWare
	// e.Use(myMw.TransactionHandler(db.Init()))

	// Routes
	v1 := test.Group("/api/v1")
	{
		// v1.POST("/search", api.PostMember())
		v1.GET("/search", api.GetList)
		// v1.GET("/members/:id", api.GetMember())
	}

	test.GET("/test", api.Test)
	test.GET("/test1", api.Test1())
	test.GET("/", func(c echo.Context) error {
        return c.String(200, "Test")
	})

	// Server
	e := echo.New()

	e.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		host := hosts[req.Host]

		if host == nil {
			err = echo.ErrNotFound
		} else {
			host.Echo.ServeHTTP(res, req)
		}
		return
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
