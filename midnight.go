package gu

// 获取前 n 天凌晨时间，东八区上，unix 时间戳取余 60*60*24 从早上8点开始
func GetMidnightBeforeDays(currentTime int64, days int64) int64 {
	var (
		sixteenHours  int64
		eightHours    int64
		leftTime      int64
		todayMidnight int64
	)

	sixteenHours = 60 * 60 * 16
	eightHours = 60 * 60 * 8
	leftTime = currentTime % (60 * 60 * 24)

	if leftTime <= sixteenHours {
		todayMidnight = currentTime - leftTime - eightHours
	} else {
		todayMidnight = currentTime - (leftTime - sixteenHours)
	}
	return todayMidnight - 60*60*24*days
}
