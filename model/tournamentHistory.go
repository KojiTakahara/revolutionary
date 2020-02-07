package model

import "time"

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
	Format     string
	Date       time.Time
}
