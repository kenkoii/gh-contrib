package ghcontrib

import (
	"net/http"

	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
)

func init() {
	e := echo.New()
	e.Use(mw.Logger())

	e.GET("/", mainHandler)
	http.HandleFunc("/", e.ServeHTTP)
}

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}
