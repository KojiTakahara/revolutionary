package api

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
)

const url = "https://dm.takaratomy.co.jp/card/"

func GetAndRegistRace(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	client := urlfetch.Client(ctx)
	resp, _ := client.Get(url)
	doc, _ := goquery.NewDocumentFromResponse(resp)
	races := []string{}
	doc.Find("#race > option").Each(func(_ int, s *goquery.Selection) {
		race := s.Text()
		if race != "全ての項目から検索する" {
			result, _ := GetRaceByName(ctx, race)
			if result == nil {
				RegistRace(ctx, race)
			}
			races = append(races, race)
		}
	})
	return c.JSON(http.StatusOK, races)
}
