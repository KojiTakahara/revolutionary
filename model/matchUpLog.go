package model

import "time"

type MatchUpLog struct {
	PlayerName       string `datastore:",noindex"`
	PlayerId         string
	PlayerType       string
	PlayerRace       string
	PlayerLight      bool
	PlayerWater      bool
	PlayerDark       bool
	PlayerFire       bool
	PlayerNature     bool
	PlayerZero       bool
	PlayerUseCards   []string `datastore:",noindex"`
	OpponentName     string   `datastore:",noindex"`
	OpponentId       string
	OpponentType     string
	OpponentRace     string
	OpponentLight    bool
	OpponentWater    bool
	OpponentDark     bool
	OpponentFire     bool
	OpponentNature   bool
	OpponentZero     bool
	OpponentUseCards []string `datastore:",noindex"`
	Url              string   `datastore:",noindex"`
	Format           string
	Win              bool
	Date             time.Time
}
