package model

import "time"

type Tournament struct {
	Format       string
	Participants int
	Date         time.Time
}
