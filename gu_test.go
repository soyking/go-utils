package gu

import (
	"testing"
	"time"
)

func TestStructToKey(t *testing.T) {
	type Request struct {
		BeginDate   string
		EndDate     string
		AccountList []string
		Numbers     []int
	}
	request := Request{}
	request.BeginDate = "2016-03-01"
	request.EndDate = "2016-03-07"
	request.AccountList = []string{"eb12d89d-fecf-4bba-9396-94b831ce3ee3"}
	request.Numbers = []int{3, 4, 5}

	t.Log(StructToKey(&request))
}

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
