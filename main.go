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

var url = `https://api.github.com/repos/%s/%s/commits?
			access_token=%s
			&since=%s
			&until=%s
			&author=%s`

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

type ContribRequest struct {
	UserOrg     string
	Repo        string
	Author      string
	DateSince   string
	DateUntil   string
	AccessToken string
}

func fetchGithubContribs(ContribRequest) {

}
