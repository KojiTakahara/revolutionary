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
