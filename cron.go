package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
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

func replaceRace(histories []TournamentHistory, c appengine.Context) {
	for i := range histories {
		history := histories[i]
		if history.Race != "" && 0 <= strings.Index(history.Race, "..") {
			race := Race{}
			key := datastore.NewKey(c, "Race", history.Race, 0, nil)
			if err := datastore.Get(c, key, &race); err != nil {
				c.Warningf(err.Error() + " : " + history.Race)
			} else if history.Race != race.TrueName {
				history.Race = race.TrueName
				_, err := updateTournamentHistory(history, c)
				if err != nil {
					c.Criticalf(err.Error())
				}
			}
		}
	}
}

/**
 * 種族を書き換える
 */
func FormatRace(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	days, _ := strconv.Atoi(params["days"])
	t := now().AddDate(0, 0, days) // x日前
	histories := getTournamentHistoryByDate(t, c)
	replaceRace(histories, c)
	r.JSON(200, t)
}

/**
 * デッキタイプを登録する
 */
func ChangeType(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	//days, _ := strconv.Atoi(params["days"])
	//t := now().AddDate(0, 0, days) // x日前
	//histories := getTournamentHistoryByDate(t, c)
	q := datastore.NewQuery("TournamentHistory")
	q = q.Filter("Type =", "")
	histories := make([]TournamentHistory, 0, 10)
	_, err := q.GetAll(c, &histories)
	if err != nil {
		c.Criticalf(err.Error())
	}
	for i := range histories {
		history := histories[i]
		q = datastore.NewQuery("DeckType")
		q = q.Filter("Light =", history.Light)
		q = q.Filter("Water =", history.Water)
		q = q.Filter("Dark =", history.Dark)
		q = q.Filter("Fire =", history.Fire)
		q = q.Filter("Nature =", history.Nature)
		q = q.Filter("Zero =", history.Zero)
		if history.Race != "" {
			q = q.Filter("Race =", history.Race)
		}
		if history.Type != "" {
			q = q.Filter("Type =", history.Type)
		}
		types := make([]DeckType, 0, 10)
		q.GetAll(c, &types)
		if 0 < len(types) {
			history.Type = types[0].TrueType
			updateTournamentHistory(history, c)
		}
	}
	r.JSON(200, histories)
}

/**
 * 更新処理
 */
func updateTournamentHistory(history TournamentHistory, c appengine.Context) (k *datastore.Key, err error) {
	h := &TournamentHistory{}
	h.Id = history.Id
	h.PlayerName = history.PlayerName
	h.PlayerId = history.PlayerId
	h.Type = history.Type
	h.Race = history.Race
	h.Light = history.Light
	h.Water = history.Water
	h.Dark = history.Dark
	h.Fire = history.Fire
	h.Nature = history.Nature
	h.Zero = history.Zero
	h.Win = history.Win
	h.Date = history.Date
	keyStr := h.Date.Format("20060102") + "_" + history.PlayerId
	key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
	return datastore.Put(c, key, h)
}

/**
 * 日付に一致するTournamentHistoryを取得
 */
func getTournamentHistoryByDate(t time.Time, c appengine.Context) []TournamentHistory {
	dateStr := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
	start, _ := time.Parse("2006-1-2", dateStr)
	end, _ := time.Parse("2006-1-2 15:04:05", dateStr+" 23:59:59")
	q := datastore.NewQuery("TournamentHistory")
	q = q.Filter("Date >=", start)
	q = q.Filter("Date <=", end)
	histories := make([]TournamentHistory, 0, 10)
	_, err := q.GetAll(c, &histories)
	if err != nil {
		c.Criticalf(err.Error())
	}
	return histories
}

func CreateTournamentHistory(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
	tournamentId, _ := strconv.Atoi(params["tournamentId"])
	if tournamentId <= 0 {
		t, _ := time.Parse("2006-01-02", "2007-04-11")
		d := now().Sub(t)
		tournamentId = int(math.Floor(d.Hours()/24)) + tournamentId
	}
	scrapingVault(tournamentId, req)
	//replaceRace(histories, appengine.NewContext(req))
	r.JSON(200, "")
}

func scrapingVault(id int, req *http.Request) {
	c := appengine.NewContext(req)
	//histories := make([]TournamentHistory, 0, 10)
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
			playerLink, _ := a.Attr("href")
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
			history.Id = id
			history.PlayerName = playerName
			history.Date = date
			history.PlayerId = strings.Trim(playerLink, "/author/")
			history.Type = deckTypes[0]
			history.Win = countWin(playerName, winPlayers)
			keyStr := date.Format("20060102") + "_" + history.PlayerId
			key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
			_, err := datastore.Put(c, key, history)
			if err != nil {
				c.Criticalf("save error. " + keyStr)
			}
		}
	})
}

/**
 * 開催日付を取得
 */
func getDate(doc *goquery.Document) time.Time {
	info := doc.Find("#rightContainer p").Text()
	runes := []rune(info)
	date := time.Now()
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
