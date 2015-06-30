package revolutionary

import (
	"time"
)

type TournamentHistory struct {
	Id         int
	PlayerName string `datastore:",noindex"`
	PlayerId   string
	Type       string
	Race       string
	Light      bool
	Water      bool
	Dark       bool
	Fire       bool
	Nature     bool
	Zero       bool
	Win        int
	Date       time.Time
}

type Race struct {
	TrueName string
}

type DeckType struct {
	Type   string
	Race   string
	Light  bool
	Water  bool
	Dark   bool
	Fire   bool
	Nature bool
	Zero   bool
}
