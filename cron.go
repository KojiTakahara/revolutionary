package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const vaultUrl = "http://dmvault.ath.cx/duel/tournament_history.php?tournamentId="

func CreateTournamentHistory(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	tournamentId, _ := strconv.Atoi(params["tournamentId"])
	if tournamentId <= 0 {
		t, _ := time.Parse("2006-01-02", "2007-04-11")
		d := now().Sub(t)
		tournamentId = int(math.Floor(d.Hours()/24)) + tournamentId
	}
	scrapingVault(tournamentId, req)
	r.JSON(200, "")
}

func scrapingVault(id int, req *http.Request) {
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	resp, _ := client.Get(vaultUrl + strconv.Itoa(id))
	doc, _ := goquery.NewDocumentFromResponse(resp)

	date := getDate(doc)

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
			history.Date = date
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

			keyStr := date.Format("20060102") + "_" + history.PlayerId
			key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
			_, err := datastore.Put(c, key, history)
			if err != nil {
				c.Criticalf("save error. " + keyStr)
			}
			//datastore.Delete(c, key)
		}
	})
}

/**
 * 開催日付を取得
 */
func getDate(doc *goquery.Document6) time.Time {
	info := doc.Find("#rightContainer p").Text()
	runes := []rune(info)
	date := now()
	switch len(info) {
	case 94:
		date, _ = stringToTime(string(runes[5:20]))
	case 95:
		date, _ = stringToTime(string(runes[5:21]))
	case 96:
		date, _ = stringToTime(string(runes[5:22]))
	}
	return date
}

/**
 * 勝利数を取得
 */
func countWin(p string, winP []string) int {
	result := 0
	for i := range winP {
		if p == winP[i] {
			result++
		}
	}
	return result
}
