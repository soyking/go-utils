package gu

import (
	"time"
)

func ParseYYYYMMDD(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}

// 取区间内连续时间
func GetConsequentTime(beginDate, endDate string) ([]string, error) {
	bt, err := ParseYYYYMMDD(beginDate)
	if err != nil {
		return nil, err
	}

	et, err := ParseYYYYMMDD(endDate)
	if err != nil {
		return nil, err
	}

	var ct []string
	for bt.Unix() <= et.Unix() {
		ct = append(ct, bt.Format("2006-01-02"))
		bt = bt.AddDate(0, 0, 1)
	}

	return ct, nil
}
