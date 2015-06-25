package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"time"
)

func GetTournamentHistory(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	u, _ := url.Parse(req.URL.String())
	params := u.Query()
	q := datastore.NewQuery("TournamentHistory")
	if len(params["startDate"]) != 0 {
		start, _ := time.Parse("2006-01-02", params["startDate"][0])
		q = q.Filter("Date >=", start)
	}
	if len(params["endDate"]) != 0 {
		end, _ := time.Parse("2006-01-02 15:04:05", params["endDate"][0]+" 23:59:59")
		q = q.Filter("Date <=", end)
	}
	histories := make([]TournamentHistory, 0, 10)
	_, err := q.GetAll(c, &histories)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
		return
	}
	r.JSON(200, histories)
}
