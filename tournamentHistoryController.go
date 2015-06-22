package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
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
	doc, _ := goquery.NewDocumentFromResponse(resp)
	doc.Find(".player").Each(func(_ int, s *goquery.Selection) {
		player := s.Find("a").Text()
		if player != "" {
			c.Infof("%s", player)

			color := s.Find(".civilcube").Text()
			c.Infof("%s", color)

			deckType := s.Find(".fontS").Text()
			c.Infof("%s", deckType)
			c.Infof("========")
		}
	})

	loop := true
	gameCount := 1
	for loop {
		p := doc.Find("#game_" + strconv.Itoa(gameCount) + " div").Text()
		c.Infof("%s", p)
		if p == "" {
			loop = false
		}
		gameCount++
	}
	result := &TournamentHistory{}
	result.Date = now()
	return result
}
