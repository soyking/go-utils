package gu

import "time"

// 获取前 n 天凌晨时间
func GetMidnightBeforeDays(t time.Time, days int) time.Time {
	leftTime :=
		(time.Duration(days*24+t.Hour()))*time.Hour +
			time.Duration(t.Minute())*time.Minute +
			time.Duration(t.Second())*time.Second +
			time.Duration(t.Nanosecond())
	return t.Add(-leftTime)
}
