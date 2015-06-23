package revolutionary

import (
	"time"
)

type TournamentHistory struct {
	PlayerName string
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
