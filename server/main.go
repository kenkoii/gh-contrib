package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

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
	authors := strings.Split(c.QueryParam("author"), ",")

	var contribReqs []*models.ContribRequest
	for _, author := range authors {
		contribReq := &models.ContribRequest{
			UserOrg:   c.QueryParam("userOrg"),
			Repo:      c.QueryParam("repo"),
			Author:    author,
			DateSince: c.QueryParam("since"),
			DateUntil: c.QueryParam("until"),
		}
		contribReqs = append(contribReqs, contribReq)
	}

	contribs := make(map[string][]*models.DailyContrib)

	for _, contribReq := range contribReqs {
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

		for i := 0; i < len(gcr); i++ {
			date := fmt.Sprintf("%s", gcr[i].Commit.Author.Date.Format("2006-01-02"))

			var found bool

			for _, v := range contribs[date] {
				if v.Author == gcr[i].Commit.Author.Name {
					v.Commits++
					found = true
				}
			}
			if !found {
				dc := &models.DailyContrib{
					Author:  gcr[i].Commit.Author.Name,
					Commits: 1,
				}
				contribs[date] = append(contribs[date], dc)
			}
		}
	}

	//Do something with gcr
	return c.JSON(http.StatusOK, contribs)
}

func fetchGithubContribs(ctx context.Context, cr *models.ContribRequest) ([]models.GithubCommitsResponse, error) {
	url := fmt.Sprintf(baseUrl, cr.UserOrg, cr.Repo, cr.DateSince, cr.DateUntil, cr.Author, os.Getenv("GH_TOKEN"))
	log.Infof(ctx, "Url: %s", url)
	res, err := urlfetch.Client(ctx).Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	var gcr []models.GithubCommitsResponse

	if err := json.NewDecoder(res.Body).Decode(&gcr); err != nil {
		return nil, err
	}

	return gcr, nil
}
