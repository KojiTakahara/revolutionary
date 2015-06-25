package revolutionary

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const location = "Asia/Tokyo"

/**
 * 現在時刻を取得する
 */
func now() time.Time {
	return time.Now().Local().In(time.FixedZone(location, 9*60*60))
}

/**
 * 日付文字列をTimeに変える
 * @param 日付文字列(yyyy年MM月dd日hh時mm分)
 */
func stringToTime(s string) (t time.Time, err error) {
	s = strings.Replace(s, "年", "_", -1)
	s = strings.Replace(s, "月", "_", -1)
	s = strings.Replace(s, "日", "_", -1)
	s = strings.Replace(s, "時", "_", -1)
	s = strings.Replace(s, "分", "", -1)
	v := strings.Split(s, "_")
	value := ""
	for i := range v {
		j, _ := strconv.Atoi(v[i])
		z := fmt.Sprintf("%02d", j)
		value += z
		switch i {
		case 0:
			value += "-"
		case 1:
			value += "-"
		case 2:
			value += " "
		case 3:
			value += ":"
		}
	}
	return time.Parse("2006-01-02 15:04", value)
}
