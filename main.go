package ghcontrib

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/kenkoii/a-fis-gh-contrib/models"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	//Require all fields
	govalidator.SetFieldsRequiredByDefault(true)
	e := echo.New()
	e.Use(mw.Logger())

	e.GET("/", mainHandler)
	http.HandleFunc("/", e.ServeHTTP)
}

var baseUrl = `https://api.github.com/repos/%s/%s/commits?
			access_token=%s
			&since=%s
			&until=%s
			&author=%s`

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func getContrib(c echo.Context) error {
	var contribReq models.ContribRequest

	if err := c.Bind(&contribReq); err != nil {
		return err
	}

	valid, err := govalidator.ValidateStruct(contribReq)
	if err != nil {
		return err
	}

	if valid == false {
		return fmt.Errorf("Struct is invalid")
	}
	//Call Fetch req
	gcr, err := fetchGithubContribs(c.Request().Context(), contribReq)
	if err != nil {
		return err
	}

	//Do something with gcr
	return c.JSON(http.StatusOK, gcr)
}

func fetchGithubContribs(c context.Context, cr models.ContribRequest) (models.GithubCommitsResponse, error) {
	url := fmt.Sprintf(baseUrl, cr.UserOrg, os.Getenv("GH_TOKEN"), cr.DateSince, cr.DateUntil, cr.Author)
	res, err := urlfetch.Client(c).Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var gcr models.GithubCommitsResponse

	if err := json.NewDecoder(res.Body).Decode(&gcr); err != nil {
		return nil, err
	}

	return gcr, nil
}
