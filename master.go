package revolutionary

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"net/http"
)

func CreateRaceData(r render.Render, req *http.Request) {
	races := map[string]string{
		"スノーフェ..":       "スノーフェアリー",
		"エンジェ..":        "エンジェル・コマンド",
		"ガイアール・コ..":     "ガイアール・コマンド・ドラゴン",
		"ガイアール・コマンド・..": "ガイアール・コマンド・ドラゴン",
		"スプラ..":         "スプラッシュ・クイーン",
		"ヒューマノイド..":     "ヒューマノイド爆",
		"リキッド・ピープル..":   "リキッド・ピープル閃",
	}
	for key, value := range races {
		c := appengine.NewContext(req)
		race := &Race{}
		race.TrueName = value
		keyStr := key
		key := datastore.NewKey(c, "Race", keyStr, 0, nil)
		_, err := datastore.Put(c, key, race)
		if err != nil {
			c.Criticalf("save error. " + keyStr)
		}
	}
	r.JSON(200, "")
}
