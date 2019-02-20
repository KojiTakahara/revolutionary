package api

import (
	"net/http"

	"github.com/KojiTakahara/revolutionary/model"
	"github.com/labstack/echo"
)

func GetTournamentHistory(e echo.Context) error {
	// c := appengine.NewContext(e.Request())
	// q := datastore.NewQuery("TournamentHistory")
	// if len(e.Param("startDate")) != 0 {
	// 	start, _ := time.Parse("2006-01-02", e.Param("startDate"))
	// 	q = q.Filter("Date >=", start)
	// }
	// if len(e.Param("endDate")) != 0 {
	// 	end, _ := time.Parse("2006-01-02 15:04:05", e.Param("endDate")+" 23:59:59")
	// 	q = q.Filter("Date <=", end)
	// }
	histories := make([]model.TournamentHistory, 0, 10)
	// _, err := q.GetAll(c, &histories)
	// if err != nil {
	// 	log.Errorf(c, "%v", err)
	// 	return e.JSON(http.StatusNoContent, err)
	// }
	// if len(e.Param("count")) != 0 {
	// 	count, _ := strconv.Atoi(e.Param("count"))
	// 	newHistories := make([]model.TournamentHistory, 0, 10)
	// 	for i := range histories {
	// 		if count <= histories[i].Win {
	// 			newHistories = append(newHistories, histories[i])
	// 		}
	// 	}
	// 	return e.JSON(http.StatusOK, newHistories)
	// }
	return e.JSON(http.StatusOK, histories)
}
