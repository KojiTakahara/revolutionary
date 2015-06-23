package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func CreateTournamentHistory(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	tournamentId, _ := strconv.Atoi(params["tournamentId"])

	t, _ := time.Parse("2006-01-02", "2014-12-31")
	d := now().Sub(t)
	fmt.Println(d)

	scrapingVault(tournamentId, req)
	r.JSON(200, "")
}

func scrapingVault(id int, req *http.Request) {
	url := "http://dmvault.ath.cx/duel/tournament_history.php?tournamentId=" + strconv.Itoa(id)
	c := appengine.NewContext(req)
	now := now()
	client := urlfetch.Client(c)
	resp, _ := client.Get(url)
	doc, _ := goquery.NewDocumentFromResponse(resp)
	winPlayers := []string{}
	loop := true
	gameCount := 1
	for loop {
		p := doc.Find("#game_" + strconv.Itoa(gameCount) + " div").Text()
		winPlayers = append(winPlayers, p)
		if p == "" {
			loop = false
		}
		gameCount++
	}
	doc.Find(".player").Each(func(_ int, s *goquery.Selection) {
		a := s.Find("a")
		playerName := a.Text()

		if playerName != "" {
			history := &TournamentHistory{}
			history.PlayerName = playerName
			playerLink, _ := a.Attr("href")
			history.PlayerId = strings.Trim(playerLink, "/author/")
			history.Date = now
			s.Find(".civilcube").Each(func(_ int, s *goquery.Selection) {
				color := s.Text()
				if color == "光" {
					history.Light = true
				} else if color == "水" {
					history.Water = true
				} else if color == "闇" {
					history.Dark = true
				} else if color == "火" {
					history.Fire = true
				} else if color == "自" {
					history.Nature = true
				} else if color == "ゼ" {
					history.Zero = true
				}
			})
			deckTypes := strings.Split(s.Find(".fontS").Text(), "（")
			if 1 < len(deckTypes) {
				history.Race = strings.Trim(deckTypes[1], "）")
			}
			history.Type = deckTypes[0]
			history.Win = countWin(playerName, winPlayers)

			keyStr := now.Format("20060102") + "_" + history.PlayerId
			key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
			_, err := datastore.Put(c, key, history)
			if err != nil {
				c.Criticalf("save error. " + keyStr)
			}
		}
	})
}

func countWin(p string, winP []string) int {
	result := 0
	for i := range winP {
		if p == winP[i] {
			result++
		}
	}
	return result
}
