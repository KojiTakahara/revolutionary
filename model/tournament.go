package model

import "time"

type Tournament struct {
	Id           int
	Format       string
	Participants int
	Date         time.Time
}
