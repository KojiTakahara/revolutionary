package revolutionary

import (
	"api"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
)

var m *martini.Martini

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())
	m.Get("/api/cron/tournament_history/:tournamentId", CreateTournamentHistory)
	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}
