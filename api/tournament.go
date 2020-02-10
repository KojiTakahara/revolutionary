package api

import (
	"net/http"
	"time"

	"github.com/KojiTakahara/revolutionary/model"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func GetTournamentHistory(e echo.Context) error {
	ctx := appengine.NewContext(e.Request())
	q := datastore.NewQuery("TournamentHistory")
	if len(e.QueryParam("startDate")) != 0 {
		start, _ := time.Parse("2006-01-02", e.QueryParam("startDate"))
		q = q.Filter("Date >=", start)
	}
	if len(e.QueryParam("endDate")) != 0 {
		end, _ := time.Parse("2006-01-02 15:04:05", e.QueryParam("endDate")+" 23:59:59")
		q = q.Filter("Date <=", end)
	}
	if len(e.QueryParam("race")) != 0 {
		q = q.Filter("Race =", e.QueryParam("race"))
	}
	if len(e.QueryParam("type")) != 0 {
		q = q.Filter("Type =", e.QueryParam("type"))
	}
	histories := make([]model.TournamentHistory, 0, 10)
	_, err := q.GetAll(ctx, &histories)
	if err != nil {
		log.Errorf(ctx, "%v", err)
	}
	return e.JSON(http.StatusOK, histories)
}
