package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/KojiTakahara/revolutionary/model"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func Get(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	matchUpLog, err := GetMatchUpLogByKey(ctx, c.Param("key"))
	if err != nil {
		log.Errorf(ctx, "%v", err)
		panic(err)
	}
	return c.JSON(http.StatusOK, matchUpLog)
}

func Find(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	matchUpLogs := []model.MatchUpLog{}
	q := datastore.NewQuery("MatchUpLog")
	if len(c.QueryParam("type")) != 0 {
		q = q.Filter("PlayerType =", c.QueryParam("type"))
	}
	if len(c.QueryParam("race")) != 0 {
		q = q.Filter("PlayerRace =", c.QueryParam("race"))
	}
	if len(c.QueryParam("light")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("light"))
		q = q.Filter("PlayerLight =", b)
	}
	if len(c.QueryParam("water")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("water"))
		q = q.Filter("PlayerWater =", b)
	}
	if len(c.QueryParam("dark")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("dark"))
		q = q.Filter("PlayerDark =", b)
	}
	if len(c.QueryParam("fire")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("fire"))
		q = q.Filter("PlayerFire =", b)
	}
	if len(c.QueryParam("nature")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("nature"))
		q = q.Filter("PlayerNature =", b)
	}
	if len(c.QueryParam("zero")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("zero"))
		q = q.Filter("PlayerZero =", b)
	}
	if len(c.QueryParam("startDate")) != 0 {
		start, _ := time.Parse("2006-01-02", c.QueryParam("startDate"))
		q = q.Filter("Date >=", start)
	}
	if len(c.QueryParam("endDate")) != 0 {
		end, _ := time.Parse("2006-01-02 15:04:05", c.QueryParam("endDate")+" 23:59:59")
		q = q.Filter("Date <=", end)
	}
	if len(c.QueryParam("win")) != 0 {
		b, _ := strconv.ParseBool(c.QueryParam("win"))
		q = q.Filter("Win =", b)
	}
	if len(c.QueryParam("format")) != 0 {
		q = q.Filter("Format =", c.QueryParam("format"))
	}
	_, err := q.GetAll(ctx, &matchUpLogs)
	if err != nil {
		log.Errorf(ctx, "%v", err)
		panic(err)
	}
	return c.JSON(http.StatusOK, matchUpLogs)
}
