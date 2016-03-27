package tu

import (
	"testing"
	"time"
)

func TestMidnight(t *testing.T) {
	currentTime := time.Now().Unix()
	// 凌晨
	t0 := GetMidnightBeforeDays(currentTime, 0)
	t.Logf("%s", time.Unix(t0, 0).String())
	// 今晚凌晨
	t1 := GetMidnightBeforeDays(currentTime, -1)
	t.Logf("%s", time.Unix(t1, 0).String())
	// 昨天凌晨
	t2 := GetMidnightBeforeDays(currentTime, 1)
	t.Logf("%s", time.Unix(t2, 0).String())
}

func TestGetConsequentTime(t *testing.T) {
	dates, err := GetConsequentTime("2016-03-01", "2016-03-07")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%v", dates)
}
