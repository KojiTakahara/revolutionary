package revolutionary.api

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"strings"
)

func CreateTournamentHistory(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	tournamentId, _ := strconv.Atoi(params["tournamentId"])
	history := scrapingVault(tournamentId, req)
	c := appengine.NewContext(req)
	key := datastore.NewKey(c, "TournamentHistory", "", 0, nil)
	key, err := datastore.Put(c, key, history)
	if err != nil {
		c.Criticalf("save error.")
	} else {
		c.Infof("success.")
	}
	r.JSON(200, "")
}

func scrapingVault(id int, req *http.Request) *TournamentHistory {
	url := "http://dmvault.ath.cx/duel/tournament_history.php?tournamentId=" + strconv.Itoa(id)
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	resp, _ := client.Get(url)
	return &TournamentHistory{}
}
