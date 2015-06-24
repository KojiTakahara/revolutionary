package revolutionary

import (
	"time"
)

const location = "Asia/Tokyo"

func now() time.Time {
	return time.Now().Local().In(time.FixedZone(location, 9*60*60))
}

func stringToTime(s string) time.Time {
	return now()
}
