package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/asaskevich/govalidator"
	"github.com/kenkoii/gh-contrib/models"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

func init() {
	//Require all fields
	govalidator.SetFieldsRequiredByDefault(true)
	e := echo.New()
	e.Use(mw.Logger())

	e.GET("/", mainHandler)
	e.GET("/api/v1/commits", getContribHandler)
	http.Handle("/", e)
}

var baseUrl = `https://api.github.com/repos/%s/%s/commits?since=%s&until=%s&author=%s&access_token=%s`

func mainHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

func getContribHandler(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	contribReq := &models.ContribRequest{
		UserOrg:   c.QueryParam("userOrg"),
		Repo:      c.QueryParam("repo"),
		Author:    c.QueryParam("author"),
		DateSince: c.QueryParam("since"),
		DateUntil: c.QueryParam("until"),
	}
	// if err := c.Bind(&contribReq); err != nil {
	// 	return err
	// }

	log.Infof(ctx, "Req Body: %v", contribReq)

	valid, err := govalidator.ValidateStruct(contribReq)
	if err != nil {
		log.Errorf(ctx, "Invalid struct")
		return err
	}

	if valid == false {
		log.Errorf(ctx, "Govalidator error")
		return fmt.Errorf("Struct is invalid")
	}
	//Call Fetch req
	gcr, err := fetchGithubContribs(ctx, contribReq)
	if err != nil {
		log.Errorf(ctx, "Fetch error: %s", err.Error())
		return err
	}

	//Do something with gcr
	return c.JSON(http.StatusOK, gcr)
}

func fetchGithubContribs(ctx context.Context, cr *models.ContribRequest) (models.GithubCommitsResponse, error) {
	log.Infof(ctx, "Hello World")
	url := fmt.Sprintf(baseUrl, cr.UserOrg, cr.Repo, cr.DateSince, cr.DateUntil, cr.Author, os.Getenv("GH_TOKEN"))
	log.Infof(ctx, "Url: %s", url)
	res, err := urlfetch.Client(ctx).Get(url)
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
