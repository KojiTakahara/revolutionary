package main

// import (
// 	"fmt"
// 	"net/http"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/go-martini/martini"
// 	"github.com/martini-contrib/render"

// 	"appengine"
// 	"appengine/datastore"
// )

// const vaultUrl = "http://dmvault.ath.cx/duel/tournament_history.php?tournamentId="

// func replaceRace(histories []TournamentHistory, c appengine.Context) {
// 	for i := range histories {
// 		history := histories[i]
// 		if history.Race != "" && 0 <= strings.Index(history.Race, "..") {
// 			race := Race{}
// 			key := datastore.NewKey(c, "Race", history.Race, 0, nil)
// 			if err := datastore.Get(c, key, &race); err != nil {
// 				c.Warningf(err.Error() + " : " + history.Race)
// 			} else if history.Race != race.TrueName {
// 				history.Race = race.TrueName
// 				_, err := updateTournamentHistory(history, c)
// 				if err != nil {
// 					c.Criticalf(err.Error())
// 				}
// 			}
// 		}
// 	}
// }

// /**
//  * 種族を書き換える
//  */
// func FormatRace(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
// 	c := appengine.NewContext(req)
// 	days, _ := strconv.Atoi(params["days"])
// 	t := Now().AddDate(0, 0, days) // x日前
// 	histories := getTournamentHistoryByDate(t, c)
// 	replaceRace(histories, c)
// 	r.JSON(200, t)
// }

// /**
//  * デッキタイプを登録する
//  */
// func ChangeType(r render.Render, params martini.Params, w http.ResponseWriter, req *http.Request) {
// 	c := appengine.NewContext(req)
// 	//days, _ := strconv.Atoi(params["days"])
// 	//t := Now().AddDate(0, 0, days) // x日前
// 	//histories := getTournamentHistoryByDate(t, c)
// 	q := datastore.NewQuery("TournamentHistory")
// 	q = q.Filter("Type =", "")
// 	histories := make([]TournamentHistory, 0, 10)
// 	_, err := q.GetAll(c, &histories)
// 	if err != nil {
// 		c.Criticalf(err.Error())
// 	}
// 	for i := range histories {
// 		history := histories[i]
// 		q = datastore.NewQuery("DeckType")
// 		q = q.Filter("Light =", history.Light)
// 		q = q.Filter("Water =", history.Water)
// 		q = q.Filter("Dark =", history.Dark)
// 		q = q.Filter("Fire =", history.Fire)
// 		q = q.Filter("Nature =", history.Nature)
// 		q = q.Filter("Zero =", history.Zero)
// 		if history.Race != "" {
// 			q = q.Filter("Race =", history.Race)
// 		}
// 		if history.Type != "" {
// 			q = q.Filter("Type =", history.Type)
// 		}
// 		types := make([]DeckType, 0, 10)
// 		q.GetAll(c, &types)
// 		if 0 < len(types) {
// 			history.Type = types[0].TrueType
// 			updateTournamentHistory(history, c)
// 		}
// 	}
// 	r.JSON(200, histories)
// }

// /**
//  * 更新処理
//  */
// func updateTournamentHistory(history TournamentHistory, c appengine.Context) (k *datastore.Key, err error) {
// 	h := &TournamentHistory{}
// 	h.Id = history.Id
// 	h.PlayerName = history.PlayerName
// 	h.PlayerId = history.PlayerId
// 	h.Type = history.Type
// 	h.Race = history.Race
// 	h.Light = history.Light
// 	h.Water = history.Water
// 	h.Dark = history.Dark
// 	h.Fire = history.Fire
// 	h.Nature = history.Nature
// 	h.Zero = history.Zero
// 	h.Win = history.Win
// 	h.Date = history.Date
// 	keyStr := h.Date.Format("20060102") + "_" + history.PlayerId
// 	key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
// 	return datastore.Put(c, key, h)
// }

// /**
//  * 日付に一致するTournamentHistoryを取得
//  */
// func getTournamentHistoryByDate(t time.Time, c appengine.Context) []TournamentHistory {
// 	dateStr := fmt.Sprintf("%d-%d-%d", t.Year(), t.Month(), t.Day())
// 	start, _ := time.Parse("2006-1-2", dateStr)
// 	end, _ := time.Parse("2006-1-2 15:04:05", dateStr+" 23:59:59")
// 	q := datastore.NewQuery("TournamentHistory")
// 	q = q.Filter("Date >=", start)
// 	q = q.Filter("Date <=", end)
// 	histories := make([]TournamentHistory, 0, 10)
// 	_, err := q.GetAll(c, &histories)
// 	if err != nil {
// 		c.Criticalf(err.Error())
// 	}
// 	return histories
// }
