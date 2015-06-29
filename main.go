package revolutionary

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

var m *martini.Martini

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/cron/tournament_history/:tournamentId", CreateTournamentHistory)
	m.Get("/cron/race/format/:days", FormatRace)
	m.Get("/api/tournament_history", GetTournamentHistory)
	// master data
	m.Get("/api/master/race", CreateRaceData)
	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}
