package cron

import (
	"context"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/KojiTakahara/revolutionary/model"
	"github.com/KojiTakahara/revolutionary/util"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/echo"
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/urlfetch"
)

const vaultUrl = "https://dmvault.ath.cx"
const tournamentUrl = "/duel/tournament_history.php?tournamentId="

/**
 * @Param tournamentId path int false "TournamentID"
 */
func CreateTournamentHistory(c echo.Context) error {
	ctx := appengine.NewContext(c.Request())
	tournamentId, _ := strconv.Atoi(c.Param("tournamentId"))
	if tournamentId <= 0 {
		tournamentId = getLatestTournamentId(tournamentId)
	}
	log.Infof(ctx, "CreateTournamentHistory tournamentId = %v", tournamentId)
	histories := scrapingVault(tournamentId, c)
	return c.JSON(http.StatusOK, histories)
}

func scrapingVault(id int, c echo.Context) []*model.TournamentHistory {
	ctx := appengine.NewContext(c.Request())
	histories := []*model.TournamentHistory{}
	client := urlfetch.Client(ctx)
	resp, _ := client.Get(vaultUrl + tournamentUrl + strconv.Itoa(id))
	doc, _ := goquery.NewDocumentFromResponse(resp)
	date := getDate(doc)
	tournamentInfo := doc.Find("#rightContainer > p").Text()
	log.Infof(ctx, tournamentInfo)
	formatName := getFormat(tournamentInfo)
	winPlayers := []string{}
	loop := true
	gameCount := 1
	for loop {
		createMatchUpLog(ctx, doc, gameCount, date, formatName)
		p := doc.Find("#game_" + strconv.Itoa(gameCount) + " div").Text()
		winPlayers = append(winPlayers, p)
		if p == "" {
			loop = false
		}
		gameCount++
	}
	playerNames := []string{}
	doc.Find(".player").Each(func(_ int, s *goquery.Selection) {
		a := s.Find("a")
		playerName := a.Text()
		if playerName != "" {
			playerLink, _ := a.Attr("href")
			deckTypes := strings.Split(s.Find(".fontS").Text(), "（")
			history := &model.TournamentHistory{
				Id:         id,
				PlayerName: playerName,
				PlayerId:   strings.Trim(playerLink, "/author/"),
				Date:       date,
				Win:        countWin(playerName, winPlayers),
				Type:       fixDeckType(ctx, deckTypes[0]),
				Format:     formatName,
			}
			s.Find(".civilcube").Each(func(_ int, s *goquery.Selection) {
				color := s.Text()
				if color == "光" {
					history.Light = true
				} else if color == "水" {
					history.Water = true
				} else if color == "闇" {
					history.Dark = true
				} else if color == "火" {
					history.Fire = true
				} else if color == "自" {
					history.Nature = true
				} else if color == "ゼ" {
					history.Zero = true
				}
			})
			if 1 < len(deckTypes) {
				history.Race = fixRace(ctx, strings.Trim(deckTypes[1], "）"))
			}
			keyStr := date.Format("20060102") + "_" + history.PlayerId
			key := datastore.NewKey(ctx, "TournamentHistory", keyStr, 0, nil)
			_, err := datastore.Put(ctx, key, history)
			if err != nil {
				log.Errorf(ctx, "keyStr: %v", err)
			}
			histories = append(histories, history)
			if !contains(playerNames, playerName) {
				playerNames = append(playerNames, playerName)
			}
		}
	})
	createTournament(ctx, id, formatName, playerNames, date)
	return histories
}

/**
 * TournamentIDを取得する
 * @Param daysAgo int true "何日前か"
 */
func getLatestTournamentId(daysAgo int) int {
	t, _ := time.Parse("2006-01-02", "2007-04-16")
	d := util.Now().Sub(t)
	return int(math.Floor(d.Hours()/24)) + daysAgo
}

/**
 * 開催日付を取得
 */
func getDate(doc *goquery.Document) time.Time {
	info := doc.Find("#rightContainer p").Text()
	runes := []rune(info)
	date := time.Now()
	switch len(info) {
	case 94:
		date, _ = util.StringToTime(string(runes[5:20]))
	case 95:
		date, _ = util.StringToTime(string(runes[5:21]))
	case 96:
		date, _ = util.StringToTime(string(runes[5:22]))
	}
	return date
}

/**
 * 勝利数を取得
 */
func countWin(p string, winP []string) int {
	result := 0
	for i := range winP {
		if p == winP[i] {
			result++
		}
	}
	return result
}

/**
 * 大会形式を取得する
 */
func getFormat(tournamentInfo string) string {
	return strings.Split(strings.Split(tournamentInfo, "／")[1], "：")[1]
}

/**
 * デッキタイプ名を整える
 * 存在しないデッキタイプを新規追加する
 */
func fixDeckType(ctx context.Context, deckType string) string {
	if len(deckType) == 0 {
		return deckType
	}
	if strings.Contains(deckType, "..") {
		// TODO
		log.Infof(ctx, "デッキタイプ含む: %v", deckType)
	} else {
		deckTypes := []model.DeckType{}
		q := datastore.NewQuery("DeckType")
		q = q.Filter("Name =", deckType)
		_, err := q.GetAll(ctx, &deckTypes)
		if err != nil {
			panic(err)
		}
		if len(deckTypes) == 0 {
			model := &model.DeckType{
				Name: deckType,
			}
			key := datastore.NewKey(ctx, "DeckType", "", 0, nil)
			datastore.Put(ctx, key, model)
		}
	}
	return deckType
}

/**
 * 種族名を整える
 * 存在しない種族を新規追加する
 */
func fixRace(ctx context.Context, race string) string {
	if len(race) == 0 {
		return race
	}
	q := datastore.NewQuery("Race")
	races := []model.Race{}
	if strings.Contains(race, "..") {
		race = strings.Replace(race, "..", "", -1)
		_, err := q.Filter("Name>=", race).Filter("Name<", race+"\ufffd").GetAll(ctx, &races)
		if err != nil {
			panic(err)
		}
		if len(races) != 0 {
			return races[0].Name
		}
	} else {
		q = q.Filter("Name =", race)
		_, err := q.GetAll(ctx, &races)
		if err != nil {
			panic(err)
		}
		if len(races) == 0 {
			model := &model.Race{
				Name: race,
			}
			key := datastore.NewKey(ctx, "Race", "", 0, nil)
			datastore.Put(ctx, key, model)
		}
	}
	return race
}

/**
 * matchUpLogを登録する
 */
func createMatchUpLog(ctx context.Context, doc *goquery.Document, gameCount int, date time.Time, format string) {
	clickEvent, _ := doc.Find("#game_" + strconv.Itoa(gameCount) + " input").Attr("onclick")
	if strings.Contains(clickEvent, "window.open") {
		winName := doc.Find("#game_" + strconv.Itoa(gameCount) + " .player").Text()
		url := vaultUrl + strings.Split(clickEvent, "'")[1]

		client := urlfetch.Client(ctx)
		resp, _ := client.Get(url)
		matchUpDoc, _ := goquery.NewDocumentFromResponse(resp)

		players := map[string]string{}
		matchUpDoc.Find(".col-sm-5").Each(func(_ int, s *goquery.Selection) {
			a := s.Find("a").First()
			playerName := a.Text()
			playerLink, _ := a.Attr("href")
			playerId := strings.Trim(playerLink, "/author/")
			if playerName == winName {
				players["winName"] = playerName
				players["winId"] = playerId
			} else {
				players["loseName"] = playerName
				players["loseId"] = playerId
			}
		})

		loserCards := []string{}
		winnerCards := []string{}
		if strings.Contains(matchUpDoc.Find(".m_l0tap").Text(), players["winName"]) {
			winnerCards = extractOnlyUsedCards(matchUpDoc, ".m_l0", winnerCards)
			loserCards = extractOnlyUsedCards(matchUpDoc, ".m_l1", loserCards)
		} else {
			loserCards = extractOnlyUsedCards(matchUpDoc, ".m_l0", loserCards)
			winnerCards = extractOnlyUsedCards(matchUpDoc, ".m_l1", winnerCards)
		}

		model1 := &model.MatchUpLog{
			PlayerName:       players["winName"],
			PlayerId:         players["winId"],
			PlayerUseCards:   winnerCards,
			OpponentName:     players["loseName"],
			OpponentId:       players["loseId"],
			OpponentUseCards: loserCards,
			Url:              url,
			Format:           format,
			Win:              true,
			Date:             date,
		}
		keyStr1 := strconv.Itoa(parseLeadingInt(clickEvent)) + "_" + players["winId"]
		putMatchUpLog(ctx, keyStr1, model1)
		model2 := &model.MatchUpLog{
			PlayerName:       players["loseName"],
			PlayerId:         players["loseId"],
			PlayerUseCards:   loserCards,
			OpponentName:     players["winName"],
			OpponentId:       players["winId"],
			OpponentUseCards: winnerCards,
			Url:              url,
			Format:           format,
			Win:              false,
			Date:             date,
		}
		keyStr2 := strconv.Itoa(parseLeadingInt(clickEvent)) + "_" + players["loseId"]
		putMatchUpLog(ctx, keyStr2, model2)
	}
}

func extractOnlyUsedCards(doc *goquery.Document, class string, list []string) []string {
	doc.Find(class).Each(func(_ int, s *goquery.Selection) {
		if !containsNgWord(s.Text()) {
			s.Find("a").Each(func(_ int, a *goquery.Selection) {
				if !contains(list, a.Text()) {
					list = append(list, a.Text())
				}
			})
		}
	})
	return list
}

func containsNgWord(s string) bool {
	if strings.Contains(s, "相手の") {
		return true
	}
	if strings.Contains(s, "vs") {
		return true
	}
	if strings.Contains(s, "攻撃") {
		return true
	}
	return false
}

func putMatchUpLog(ctx context.Context, keyStr string, model *model.MatchUpLog) {
	key := datastore.NewKey(ctx, "MatchUpLog", keyStr, 0, nil)
	_, err := datastore.Put(ctx, key, model)
	if err != nil {
		log.Errorf(ctx, "keyStr: %v", err)
	}
}

func createTournament(ctx context.Context, id int, format string, participants []string, date time.Time) {
	keyStr := date.Format("20060102")
	key := datastore.NewKey(ctx, "Tournament", keyStr, 0, nil)
	model := model.Tournament{
		Id:           id,
		Format:       format,
		Participants: len(participants),
		Date:         date,
	}
	_, err := datastore.Put(ctx, key, &model)
	if err != nil {
		log.Errorf(ctx, "error: %v", err)
	}
}

// 配列内に指定した文字列が存在するかどうか
func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

/**
 * 数字部分を抜き出す
 */
func parseLeadingInt(s string) int {
	rex := regexp.MustCompile(`\d+`).Copy()
	value, _ := strconv.Atoi(rex.FindString(s))
	return value
}
