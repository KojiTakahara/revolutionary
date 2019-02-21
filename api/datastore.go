package api

import (
	"context"

	"github.com/KojiTakahara/revolutionary/model"
	"google.golang.org/appengine/datastore"
)

func GetMatchUpLogByKey(ctx context.Context, keyStr string) (model.MatchUpLog, error) {
	matchUpLog := model.MatchUpLog{}
	key := datastore.NewKey(ctx, "MatchUpLog", keyStr, 0, nil)
	err := datastore.Get(ctx, key, &matchUpLog)
	return matchUpLog, err
}

func GetTournamentHistoryByKey(ctx context.Context, keyStr string) (model.TournamentHistory, error) {
	tournamentHistory := model.TournamentHistory{}
	key := datastore.NewKey(ctx, "TournamentHistory", keyStr, 0, nil)
	err := datastore.Get(ctx, key, &tournamentHistory)
	return tournamentHistory, err
}

func GetRaceByName(ctx context.Context, name string) (*model.Race, error) {
	races := []model.Race{}
	q := datastore.NewQuery("Race")
	_, err := q.Filter("Name>=", name).Filter("Name<", name+"\ufffd").GetAll(ctx, &races)
	if err != nil {
		return nil, err
	}
	if len(races) != 0 {
		return &races[0], nil
	}
	return nil, nil
}

func RegistRace(ctx context.Context, name string) (*model.Race, error) {
	model := &model.Race{
		Name: name,
	}
	key := datastore.NewKey(ctx, "Race", "", 0, nil)
	_, err := datastore.Put(ctx, key, model)
	return model, err
}
