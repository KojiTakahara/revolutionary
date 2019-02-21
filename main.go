// +build !appengine,!appenginevm

package main

import (
	"net/http"

	"github.com/KojiTakahara/revolutionary/api"
	"github.com/KojiTakahara/revolutionary/cron"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/appengine"
)

func main() {
	e := echo.New()
	defer e.Close()

	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	e.Use(middleware.Static("static/public"))

	http.Handle("/", e)
	g := e.Group("/api/v1")
	g.GET("/matchupLog", api.Find)
	g.GET("/matchupLog/:key", api.Get)
	g.GET("/tournament", api.GetTournamentHistory)
	g.GET("/race", api.GetAndRegistRace)
	c := e.Group("/cron/v1")
	c.GET("/tournamentHistory/:tournamentId", cron.CreateTournamentHistory)
	c.GET("/matchUpLog/:offset", cron.AddDeckTypeInfoToMatchUpLog)

	appengine.Main()
}
