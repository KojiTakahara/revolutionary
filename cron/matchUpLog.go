package cron

import (
	"context"
	"net/http"
	"strconv"

	"github.com/KojiTakahara/revolutionary/model"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
)

func AddDeckTypeInfoToMatchUpLog(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())

	t := []model.Tournament{}
	offset, _ := strconv.Atoi(c.Param("offset"))
	_, err := datastore.NewQuery("Tournament").Order("-Date").Limit(1).Offset(offset).GetAll(ctx, &t)
	if len(t) == 0 || err != nil {
		log.Errorf(ctx, "%v", err)
		return err
	}
	date := t[0].Date

	ms := []model.MatchUpLog{}
	keys, _ := datastore.NewQuery("MatchUpLog").Filter("Date =", date).GetAll(ctx, &ms)
	for i, _ := range keys {
		player := getTournamentHistory(ctx, date.Format("20060102")+"_"+ms[i].PlayerId)
		opponent := getTournamentHistory(ctx, date.Format("20060102")+"_"+ms[i].OpponentId)
		ms[i].PlayerType = player.Type
		ms[i].PlayerRace = player.Race
		ms[i].PlayerLight = player.Light
		ms[i].PlayerWater = player.Water
		ms[i].PlayerDark = player.Dark
		ms[i].PlayerFire = player.Fire
		ms[i].PlayerNature = player.Nature
		ms[i].PlayerZero = player.Zero
		ms[i].OpponentType = opponent.Type
		ms[i].OpponentRace = opponent.Race
		ms[i].OpponentLight = opponent.Light
		ms[i].OpponentWater = opponent.Water
		ms[i].OpponentDark = opponent.Dark
		ms[i].OpponentFire = opponent.Fire
		ms[i].OpponentNature = opponent.Nature
		ms[i].OpponentZero = opponent.Zero
		// key := datastore.NewKey(c, "TournamentHistory", keyStr, 0, nil)
		// datastore.Put(ctx, key, &ms[i])
	}
	datastore.PutMulti(ctx, keys, ms)
	return c.JSON(http.StatusOK, ms)
}

func getTournamentHistory(ctx context.Context, keyStr string) model.TournamentHistory {
	th := model.TournamentHistory{}
	key := datastore.NewKey(ctx, "TournamentHistory", keyStr, 0, nil)
	if err := datastore.Get(ctx, key, &th); err != nil {
		log.Errorf(ctx, "%v. keyStr %v", err, keyStr)
	}
	return th
}
